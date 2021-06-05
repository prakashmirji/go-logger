package examples

import (
	l "github.com/prakashmirji/ez-logger/pkg"
	"github.com/sirupsen/logrus"
)

func printLogMessage(num int) {
	l.Init()
	log := l.GetLogger()
	if num == 10 {
		log.Info("Hello, you are logging Info level messages")
	} else {
		log.Error("Hello, you are logging Error level messages")
	}
}

func printLogMessageWithFields() {
	l.Init()
	log := l.GetLogger()
	log.WithFields(logrus.Fields{
		"app":  "ezctl",
		"user": "root",
	}).Info("hello message with fields")
}

func printLogMessageWithHooks() {
	l.Init()
	log := l.GetLogger()
	log.AddHook(l.NewExtraFieldHook("ezctl", "dev", "defualt"))
	log.Info("logging with Hooks enabled")
}

func printLogMessageToFile(filename string) {
	l.Init()
	l.WriteToFile(filename)
	log := l.GetLogger()
	log.AddHook(l.NewExtraFieldHook("ezctl", "dev", "defualt"))
	log.Info("logging with Hooks enabled")
}