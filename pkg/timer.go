// Copyright (c) nano Author and TFG Co. All Rights Reserved.
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

package pitaya

import (
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/timer"
)

// NewTimer returns a new Timer containing a function that will be called
// with a period specified by the duration argument. It adjusts the intervals
// for slow receivers.
// The duration d must be greater than zero; if not, NewTimer will panic.
// Stop the timer to release associated resources.
func NewTimer(interval time.Duration, fn timer.Func) *timer.Timer {
	_ = "STUB: not implemented"
	return nil
}

// NewCountTimer returns a new Timer containing a function that will be called
// with a period specified by the duration argument. After count times, timer
// will be stopped automatically, It adjusts the intervals for slow receivers.
// The duration d must be greater than zero; if not, NewCountTimer will panic.
// Stop the timer to release associated resources.
func NewCountTimer(interval time.Duration, count int, fn timer.Func) *timer.Timer {
	_ = "STUB: not implemented"
	return nil
}

// add to manager

// NewAfterTimer returns a new Timer containing a function that will be called
// after duration that specified by the duration argument.
// The duration d must be greater than zero; if not, NewAfterTimer will panic.
// Stop the timer to release associated resources.
func NewAfterTimer(duration time.Duration, fn timer.Func) *timer.Timer {
	_ = "STUB: not implemented"
	return nil
}

// NewCondTimer returns a new Timer containing a function that will be called
// when condition satisfied that specified by the condition argument.
// The duration d must be greater than zero; if not, NewCondTimer will panic.
// Stop the timer to release associated resources.
func NewCondTimer(condition timer.Condition, fn timer.Func) (*timer.Timer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetTimerPrecision set the ticker precision, and time precision can not less
// than a Millisecond, and can not change after application running. The default
// precision is time.Second
func SetTimerPrecision(precision time.Duration) { _ = "STUB: not implemented"; return }

// SetTimerBacklog set the timer created/closing channel backlog, A small backlog
// may cause the logic to be blocked when call NewTimer/NewCountTimer/timer.Stop
// in main logic gorontine.
func SetTimerBacklog(c int) { _ = "STUB: not implemented"; return }
