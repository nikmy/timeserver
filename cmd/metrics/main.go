package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/nikmy/timeserver/internal/metrics"
	"github.com/nikmy/timeserver/internal/puller"
)

func main() {
	cfg := loadConfig()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	registry := metrics.NewRegistry()
	runPuller(ctx, registry, cfg)

	err := metrics.ServeHTTP(cfg.Addr, registry)
	if err != nil {
		slog.Error(err.Error())
	}
}

func runPuller(ctx context.Context, registry metrics.Registry, cfg config) {
	p := puller.New(cfg.ObservableAddr, registry)
	go func() {
		tick := time.NewTicker(cfg.PullInterval)
		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				err := p.Pull()
				if err != nil {
					slog.Error(err.Error())
				}
			}
		}
	}()
}

type config struct {
	Addr           string
	ObservableAddr string
	PullInterval   time.Duration
}

func loadConfig() config {
	var (
		port = flag.String("port", "22001", "Port")
		srv  = flag.String("service", "", "Observable service url")
		pInt = flag.Duration("interval", time.Second*5, "Pull interval")
	)
	flag.Parse()

	return config{
		Addr:           ":" + *port,
		ObservableAddr: *srv,
		PullInterval:   *pInt,
	}
}
