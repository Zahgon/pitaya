package services

import (
	"context"

	"github.com/topfreegames/pitaya/v3/examples/demo/custom_metrics/messages"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

// Room server
type Room struct {
	component.Base
	app pitaya.Pitaya
}

// NewRoom ctor
func NewRoom(app pitaya.Pitaya) *Room { _ = "STUB: not implemented"; return nil }

// SetCounter sets custom my_counter
func (r *Room) SetCounter(
	ctx context.Context,
	arg *messages.SetCounterArg,
) (*messages.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetGauge1 sets custom my_gauge_1
func (r *Room) SetGauge1(
	ctx context.Context,
	arg *messages.SetGaugeArg,
) (*messages.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetGauge2 sets custom my_gauge_2
func (r *Room) SetGauge2(
	ctx context.Context,
	arg *messages.SetGaugeArg,
) (*messages.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSummary sets custom my_summary
func (r *Room) SetSummary(
	ctx context.Context,
	arg *messages.SetSummaryArg,
) (*messages.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
