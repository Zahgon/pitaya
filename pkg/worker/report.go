package worker

import (
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/metrics"
)

// Report sends periodic of worker reports
func Report(reporters []metrics.Reporter, period time.Duration) { _ = "STUB: not implemented"; return }

func reportJobsRetry(r metrics.Reporter, retries int64) { _ = "STUB: not implemented"; return }

func reportQueueSizes(r metrics.Reporter, queues map[string]string) {
	_ = "STUB: not implemented"
	return
}

func reportJobsTotal(r metrics.Reporter, failed, processed int) {
	_ = "STUB: not implemented"
	// "failed" and "processed" always grow up,
	// so they work as count but must be reported as gauge
	return
}

func checkReportErr(metric string, err error) { _ = "STUB: not implemented"; return }
