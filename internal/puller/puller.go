package puller

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"

	"github.com/nikmy/timeserver/internal/metrics"
)

func New(url string, registry metrics.Registry) *puller {
	return &puller{
		c: resty.New().SetBaseURL(url),
		m: registry.NewNumeric("request_count", "total number of request to /time"),
	}
}

type puller struct {
	c *resty.Client
	m metrics.Numeric
}

func (p *puller) Pull() error {
	resp, err := p.c.R().Get("/statistics")
	if err != nil {
		return fmt.Errorf("cannot get stats from app server: %w", err)
	}

	count, err := strconv.ParseInt(resp.String(), 10, 64)
	if err != nil {
		return fmt.Errorf("cannot parse stats as int: %w", err)
	}

	p.m.Set(float64(count))

	return nil
}
