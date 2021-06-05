package pkg

import (
	"testing"
)

func Test(m *testing.T) {
	Init()
	l := GetLogger()
	//l.SetLevel(logrus.DebugLevel)
	l.AddHook(NewExtraFieldHook("ezctl", "dev", "defualt"))
	log.Info("This is Info message")
	log.Error("This is Error message")
	log.Debug("This is Debug message")
	log.Warn("This is Warn message")
}