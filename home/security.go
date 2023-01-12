package home

import "time"

func fixedSleep(begin time.Time, dur time.Duration) {
	future := begin.Add(dur)

	if time.Since(future) >= 0 {
		return
	}

	sleepDur := future.Sub(begin)

	time.Sleep(sleepDur)
}
