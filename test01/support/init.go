package support

import (
	"context"
	"time"
)

var (
	ctx    = context.Background()
	client gredisapi.Client
)

func Initialize() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	conMan := conman.NewSingleHost(ctx,
		conman.Host("127.0.0.1"),
		conman.Port(6379),
		conman.DB(0),
		conman.MaxActive(10),
	)
	c, err := gredis3.NewClient(ctx, conMan, gredis3.Logger(logging.GetDefaultLogger()))
	if err != nil {
		panic(err)
	}

	client = c

	initRecentWorkers()
}
