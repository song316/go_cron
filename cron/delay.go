package cron

import "time"

type DelaySchedule struct {
	Delay time.Duration
}

//Every
func Every(duration time.Duration) DelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	return DelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds()) % time.Second,
	}
}

//Next
func (ds DelaySchedule)Next(t time.Time) time.Time {
	return t.Add(ds.Delay - time.Duration(t.Nanosecond()) * time.Nanosecond)
}
