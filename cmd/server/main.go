package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v3/log"

	"github.com/nikmy/timeserver/internal/timeserver"
)

func main() {
	app := timeserver.New(getAddr())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Run(ctx)
	if err != nil {
		log.Error(err)
	}
}

func getAddr() string {
	p := flag.String("port", "22000", "Port")
	flag.Parse()
	return ":" + *p
}
