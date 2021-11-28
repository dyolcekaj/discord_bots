package bot

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dyolcekaj/discord_bots/internal/healthcheck"
	"github.com/sirupsen/logrus"
)

const DefaultPort = 8080

const envToken = "DISCORD_BOT_TOKEN"
const envClientID = "DISCORD_CLIENT_ID"

var ErrNoToken = errors.New("no bot token provided")
var ErrNoClientID = errors.New("no client ID provided")
var ErrNotStarted = errors.New("stop called without starting")

type BotOptions struct {
	Port int
}

type Bot interface {
	Run() error
}

var _ Bot = (*bot)(nil)

type bot struct {
	token string
	ctx   context.Context

	httpSrv *http.Server
}

func New(ctx context.Context, bo BotOptions) (Bot, error) {
	token := os.Getenv(envToken)
	if len(token) == 0 {
		return nil, ErrNoToken
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", healthcheck.HandlerFunc)

	port := DefaultPort
	if bo.Port != 0 {
		port = bo.Port
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return &bot{
		token: token,
		ctx:   ctx,

		httpSrv: srv,
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

	if err := b.httpSrv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
