package pitaya

import (
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/metrics/models"
)

// CreatePrometheusReporter create a Prometheus reporter instance
func CreatePrometheusReporter(serverType string, config config.MetricsConfig, customSpecs models.CustomMetricsSpec) (*metrics.PrometheusReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateStatsdReporter create a Statsd reporter instance
func CreateStatsdReporter(serverType string, config config.MetricsConfig) (*metrics.StatsdReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
