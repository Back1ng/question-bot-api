package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var Log Logger

func New() {
	Log = Logger{logrus.New()}

	Log.SetFormatter(&logrus.JSONFormatter{})

	Log.SetOutput(os.Stdout)
}
