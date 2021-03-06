package main

import (
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/rogierlommers/logrus-redis-hook"
)

func init() {
	hook, err := logredis.NewHook("localhost",
		"my_redis_key", // key to use
		"v0",           // logstash format (v0, v1)
		"my_app_name",  // your application name
		"my_hostname",  // your hostname
		"",             // password for redis authentication, leave empty for no authentication
		6379,           // redis port
		0,              // redis database
	)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func main() {
	// when hook is injected succesfully, logs will be send to redis server
	logrus.Debug("just some debug logging...")
	logrus.Info("just some info logging...")
	logrus.Warn("just some warn logging...")
	logrus.Error("just some error logging...")
	// logrus.Fatal("just some fatal logging...") // commented out, because it will kill the app
	// logrus.Panic("just some panic logging...") // commented out, because it will kill the app

	// we also support log.WithFields()
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that"}).
		Info("additional fields are being logged as well")

	// If you want to disable writing to stdout, use setOutput
	logrus.SetOutput(ioutil.Discard)
	logrus.Info("This will only be sent to Redis")
}
