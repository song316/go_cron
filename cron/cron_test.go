package cron_test

import (
	"fmt"
	"testing"
	"time"
	"github.com/song316/go_cron/cron"
	. "github.com/smartystreets/goconvey/convey"
)

const ONE_SECOND = 1 * time.Second + 10 * time.Millisecond

func TestFuncJob_Run(t *testing.T) {
	Convey("测试启动Job", t, func() {
		cron := cron.New()
		cron.Start()
		select {
		case <-time.After(ONE_SECOND):
			t.FailNow()
		case <-stop(cron):
		}
	})
}

func TestRun(t *testing.T) {
	c := cron.New();
	c.AddFunc("1 * * * * *", func() {
		fmt.Println("每秒")
	})
	c.Start()
}

func stop(cron *cron.Cron) chan bool {
	ch := make(chan bool)
	go func() {
		cron.Stop()
		ch <- true
	}()
	return ch
}