package support

import (
)

func CleanCounter() {
	passes := 0
	for {
		gredisapi.Int64(client.Do(ctx, "ZCARD", "know:"))
	}
}
