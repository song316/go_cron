package cron_test

import (
	"testing"
	"time"
	"github.com/song316/go_cron/cron"
)

const ONE_SECOND = 1 * time.Second + 10 * time.Millisecond

func TestFuncJob_Run(t *testing.T) {
	cron := cron.New()
	cron.Start()
	select {
	case <-time.After(ONE_SECOND):
		t.FailNow()
	case <-stop(cron):

	}
}

func stop(cron *cron.Cron) chan bool {
	ch := make(chan bool)
	go func() {
		cron.Stop()
		ch <- true
	}()
	return ch
}