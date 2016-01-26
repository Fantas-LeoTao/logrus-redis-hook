package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/rogierlommers/logrus-redis-hook"
)

var log = logrus.New()

func init() {
	hook, err := logredis.NewHook("localhost", 6379, "mykey")
	if err == nil {
		log.Hooks.Add(hook)
	}
}

func main() {
	// send 1000 records to redis
	for i := 0; i < 1000; i++ {
		log.Infof("logrule, number: %d", i)
		time.Sleep(1 * time.Second)
	}
}