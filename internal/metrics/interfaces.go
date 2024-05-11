package metrics

import "github.com/prometheus/client_golang/prometheus"

type Registry interface {
	NewNumeric(name string, desc string) Numeric

	prometheus.Registerer
	prometheus.Gatherer
}

type Numeric interface {
	Set(float64)
	Add(float64)
	Sub(float64)
}
