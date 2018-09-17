package support

import (
	"fmt"
	"time"
)

const (
	recentMaxWorkers = 10
)

var (
	recentLogs = make(chan Log)
)

func LogRecent(log Log) {
	recentLogs <- log
}

func initRecentWorkers() {
	var consume = func(log Log) {
		destination := fmt.Sprintf("recent:%s", log.Severity)
		msg := fmt.Sprintf("%s - %s", time.Now().Format(TimeLayout), log.Message)

		args := [][]interface{}{
			{"LPUSH", destination, msg},
			{"LTRIM", destination, 0, 99},
		}

		client.Pipeline(ctx, args)
	}

	for i := 0; i < recentMaxWorkers; i++ {
		go func() {
			for log := range recentLogs {
				consume(log)
			}
		}()
	}
}
