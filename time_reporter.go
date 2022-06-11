package ginmdw

import (
	"time"

	"github.com/gin-gonic/gin"
)

type timeReporter struct {
	sm      syncMap[string, TimeReport]
	dur     time.Duration
	collect func(map[string]TimeReport)
	init    bool
}

func TimeReporter(interval time.Duration, f func(map[string]TimeReport)) *timeReporter {
	if interval == 0 {
		interval = 500 * time.Millisecond
	}

	reporter := &timeReporter{
		sm: syncMap[string, TimeReport]{
			m: make(map[string]TimeReport),
		},
		dur:     interval,
		collect: f,
		init:    true,
	}

	reporter.loopReport()

	return reporter
}

type TimeReport struct {
	Hits        int           `json:"hits"`
	MaxDuration time.Duration `json:"maxDur"`
	MinDuration time.Duration `json:"minDur"`
}

func (r *timeReporter) Collect() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		go r.hit(c.Request.URL.Path, time.Since(t))
	}
}

func (t *timeReporter) hit(s string, latency time.Duration) {
	defer recover()

	t.sm.GetSet(s, func(tr TimeReport) TimeReport {
		if tr.MaxDuration < latency {
			tr.MaxDuration = latency
		}
		if tr.MinDuration > latency || tr.MinDuration == 0 {
			tr.MinDuration = latency
		}
		tr.Hits++

		return tr
	})
}

func (t *timeReporter) loopReport() {
	go time.AfterFunc(t.dur, t.loopReport)

	if t.init {
		t.init = false
		return
	}

	if t.collect != nil {

		t.sm.l.Lock()
		report := t.sm.m
		t.sm.m = make(map[string]TimeReport)
		defer t.sm.l.Unlock()

		t.collect(report)
	}
}
