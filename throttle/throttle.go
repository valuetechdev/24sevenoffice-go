package throttle

import (
	"context"
	"net/http"
	"time"
)

// [Throttler] blocks when requests happen faster than actions/s
//
// wraps the [time.Ticker] in helper methods
type Throttler struct {
	t *time.Ticker
}

func New(actionsPerSecond int) *Throttler {
	t := time.NewTicker(time.Second / time.Duration(actionsPerSecond))

	return &Throttler{t}
}

// request interceptor that blocks until the ticker is ready
func (rl *Throttler) Interceptor(ctx context.Context, r *http.Request) error {
	rl.Wait()
	return nil
}

// wait until the channel stops blocking
func (rl *Throttler) Wait() {
	<-rl.t.C
}

// stop the ticker
func (rl *Throttler) Stop() {
	rl.t.Stop()
}
