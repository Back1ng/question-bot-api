package interval

import (
	"time"
)

func GetActual(seconds int) (intervals []int) {
	now := time.Now()

	startDay := time.Date(now.Year(), now.Month(), now.Day(), 5, 0, 0, 0, now.Location())
	endDay := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())

	currentSeconds := int(endDay.Sub(now).Seconds())
	daySeconds := endDay.Sub(startDay).Seconds()

	if currentSeconds < 0 || currentSeconds > int(daySeconds) {
		return
	}

	for interval := 1; interval <= 24; interval++ {
		for i := int(daySeconds) / interval; i <= int(daySeconds); i += int(daySeconds) / interval {
			if currentSeconds-seconds < i && currentSeconds > i {
				mustSend := true
				for _, tempInterval := range intervals {
					if tempInterval == interval {
						mustSend = false
					}
				}
				if mustSend {
					intervals = append(intervals, interval)
				}
			}
		}
	}

	return
}
