package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rogierlommers/logrus-redis-hook"
)

func init() {
	hook, err := logredis.NewHook("localhost", 6379, "my_redis_key", "v0", "my_app_name", "my_hostname")
	if err == nil {
		log.AddHook(hook)
	} else {
		log.Error(err)
	}
}

func main() {
	// when hook is injected succesfully, logs will be send to redis server
	log.Info("just some info logging...")

	// we also support log.WithFields()
	log.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")
}
