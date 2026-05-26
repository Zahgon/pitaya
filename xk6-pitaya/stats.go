package pitaya

import (
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type pitayaMetrics struct {
	RequestResponseTime *metrics.Metric
	TimeoutRequests     *metrics.Metric
	TagsAndMeta         *metrics.TagsAndMeta
}

// registerMetrics registers the metrics for the mqtt module in the metrics registry
func registerMetrics(vu modules.VU) (pitayaMetrics, error) {
	_ = "STUB: not implemented"
	return *new(pitayaMetrics), nil
}
