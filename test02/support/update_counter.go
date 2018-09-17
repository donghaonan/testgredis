package support

import (
	"fmt"
	"time"
)

var precisions = []int64{1, 5, 60, 300, 3600, 18000, 86400}

func UpdateCounter(name string) {
	now := time.Now().Unix()

	args := [][]interface{}{}

	for _, p := range precisions {
		pnow := int64(now/p) * p
		args = append(args, [][]interface{}{
			{"ZADD", "know:", now, 0},
			{"HINCR", fmt.Sprintf("count:%d", now), pnow},
		}...)
	}

	client.Pipeline(ctx, args)
}
