package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type Hook interface {
	Levels() []logrus.Level
	Fire(*logrus.Entry) error
}

type ExtraFieldHook struct {
	service   string
	env       string
	pid       int
	namespace string
}

func NewExtraFieldHook(service, env, namespace string) *ExtraFieldHook {
	return &ExtraFieldHook{
		service: service,
		env:     env,
		namespace: namespace,
		pid:     os.Getpid(),
	}
}

func (h *ExtraFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *ExtraFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["service"] = h.service
	entry.Data["env"] = h.env
	entry.Data["pid"] = h.pid
	entry.Data["namespace"] = h.namespace
	return nil
}

func Init() {
	log = logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		//DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.InfoLevel)

}

func SetJSONFormatter() {
	// TODO if environment is prod, set to JSON else Text formatter
	log.SetFormatter(&logrus.JSONFormatter{})
}

func WriteToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func GetLogger() *logrus.Logger {
	return log
}

func SetLogger(l *logrus.Logger) {
	log = l
}

