package main

import (
	"math/rand"
	"time"

	"github.com/donghaonan/testgredis/test01/support"
)

func main() {
	support.Initialize()

	for {
		r := rand.Intn(4)
		var log support.Log

		switch r {
		case 0:
			log = support.Log{Severity: support.INFO, Message: "login"}
		case 1:
			log = support.Log{Severity: support.DEBUG, Message: "test"}
		case 2:
			log = support.Log{Severity: support.WARNING, Message: "wrong"}
		case 3:
			log = support.Log{Severity: support.ERROR, Message: "error"}
		}

		support.LogRecent(log)

		time.Sleep(time.Microsecond * 10)
	}

}
