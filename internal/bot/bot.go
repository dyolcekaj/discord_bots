package bot

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"github.com/dyolcekaj/discord_bots/internal/healthcheck"
)

type BotOptions struct {
}

type Bot interface {
	Shutdown() <-chan struct{}
	Cancel()
}

var _ Bot = (*bot)(nil)

type bot struct {
	s      *discordgo.Session
	ctx    context.Context
	cancel func()
}

func New(bo BotOptions) Bot {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthcheck.HandlerFunc)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("http server unexpectedly exited")
		}
	}()

	go func() {
		<-ctx.Done()
		log.Info("shutting http server down")
		srv.Shutdown(context.Background())
	}()

	return &bot{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (b *bot) Shutdown() <-chan struct{} {
	return b.ctx.Done()
}

func (b *bot) Cancel() {
	b.cancel()
}
