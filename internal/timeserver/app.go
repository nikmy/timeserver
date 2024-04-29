package timeserver

import (
	"context"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v3"
)

type server struct {
	*fiber.App
	nRequests atomic.Uint64
	port      string
}

func New(port string) *server {
	s := &server{
		App: fiber.New(fiber.Config{
			UnescapePath: true,
			GETOnly:      true,
		}),
		port: port,
	}

	s.Get("/time", func(c fiber.Ctx) error {
		s.nRequests.Add(1)
		return c.JSON(TimeInfo{UTC: time.Now().UTC().Format(time.DateTime)})
	})

	s.Get("/statistics", func(c fiber.Ctx) error {
		response := strconv.FormatUint(s.nRequests.Load(), 10)
		return c.Send([]byte(response))
	})

	return s
}

func (s *server) Run(ctx context.Context) error {
	return s.Listen(s.port, fiber.ListenConfig{
		//EnablePrefork:   true,
		GracefulContext: ctx,
	})
}

type TimeInfo struct {
	UTC string `json:"utc"`
}
