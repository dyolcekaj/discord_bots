package discord

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const DefaultPort = 8080

const DefaultTokenEnv = "DISCORD_BOT_TOKEN"
const DefaultClientIDEnv = "DISCORD_CLIENT_ID"

var ErrNoToken = errors.New("no bot token provided")
var ErrNoClientID = errors.New("no client ID provided")
var ErrNotStarted = errors.New("stop called without starting")

type BotOptions struct {
	Port int

	TokenEnv    string
	ClientIDEnv string
}

type Bot interface {
	Run() error
}

var _ Bot = (*bot)(nil)

type bot struct {
	token    string
	clientId string
	commands []Command

	ctx context.Context

	httpSrv *http.Server
}

func NewBot(ctx context.Context, bo BotOptions, commands []Command) (Bot, error) {
	token := getEnv(bo.TokenEnv, DefaultTokenEnv)
	if len(token) == 0 {
		return nil, ErrNoToken
	}
	clientId := getEnv(bo.ClientIDEnv, DefaultClientIDEnv)
	if len(clientId) == 0 {
		return nil, ErrNoClientID
	}

	mux := http.NewServeMux()
	mux.Handle("/", NewHandler(commands))

	port := DefaultPort
	if bo.Port != 0 {
		port = bo.Port
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return &bot{
		token:    token,
		clientId: clientId,
		commands: commands,
		ctx:      ctx,
		httpSrv:  srv,
	}, nil
}

func (b *bot) Run() error {
	// Caller cancelling context will propagate and we need
	// only pay attention to this child's done-ness
	sigCtx, cancel := signal.NotifyContext(
		b.ctx, syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	go func() {
		<-sigCtx.Done()
		logrus.Info("shutting internal http server down")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		b.httpSrv.Shutdown(ctx)
	}()

	b.registerOrUpdateCommands()

	if err := b.httpSrv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
