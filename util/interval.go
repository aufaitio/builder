package util

import (
	"time"
)

// SetIntervalAsync with a non blocking task
func SetIntervalAsync(callback func(), interval time.Duration) chan bool {
	// Setup the ticker and the channel to signal the ending of the interval
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				go callback()
			case <-clear:
				ticker.Stop()
				return
			}

		}
	}()

	return clear
}
