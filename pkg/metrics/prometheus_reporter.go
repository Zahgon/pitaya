// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/metrics/models"
)

var (
	prometheusReporter *PrometheusReporter
	once               sync.Once
)

// PrometheusReporter reports metrics to prometheus
type PrometheusReporter struct {
	serverType            string
	game                  string
	countReportersMap     map[string]*prometheus.CounterVec
	summaryReportersMap   map[string]*prometheus.SummaryVec
	histogramReportersMap map[string]*prometheus.HistogramVec
	gaugeReportersMap     map[string]*prometheus.GaugeVec
	additionalLabels      map[string]string
}

func (p *PrometheusReporter) registerCustomMetrics(
	constLabels map[string]string,
	additionalLabelsKeys []string,
	spec *models.CustomMetricsSpec,
) {
	_ = "STUB: not implemented"
	return
}

func (p *PrometheusReporter) registerMetrics(
	constLabels, additionalLabels map[string]string,
	spec *models.CustomMetricsSpec,
) {
	_ = "STUB: not implemented"
	return
}

// HandlerResponseTimeMs summary

// ProcessDelay summary

// ConnectedClients gauge

// GetPrometheusReporter gets the prometheus reporter singleton
func GetPrometheusReporter(
	serverType string,
	config config.MetricsConfig,
	metricsSpecs models.CustomMetricsSpec,
) (*PrometheusReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPrometheusReporter(
	serverType string,
	config config.MetricsConfig,
	metricsSpecs *models.CustomMetricsSpec,
) (*PrometheusReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportSummary reports a summary metric
func (p *PrometheusReporter) ReportSummary(metric string, labels map[string]string, value float64) error {
	_ = "STUB: not implemented"
	return nil
}

// ReportHistogram reports a histogram metric
func (p *PrometheusReporter) ReportHistogram(metric string, labels map[string]string, value float64) error {
	_ = "STUB: not implemented"
	return nil
}

// ReportCount reports a summary metric
func (p *PrometheusReporter) ReportCount(metric string, labels map[string]string, count float64) error {
	_ = "STUB: not implemented"
	return nil
}

// ReportGauge reports a gauge metric
func (p *PrometheusReporter) ReportGauge(metric string, labels map[string]string, value float64) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureLabels checks if labels contains the additionalLabels values,
// otherwise adds them with the default values
func (p *PrometheusReporter) ensureLabels(labels map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
