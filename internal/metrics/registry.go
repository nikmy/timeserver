package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func NewRegistry() Registry {
	return &registry{Registry: prometheus.NewRegistry()}
}

type registry struct {
	*prometheus.Registry
}

func (r registry) NewNumeric(name string, desc string) Numeric {
	c := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: desc,
	})

	r.MustRegister(c)

	return c
}
