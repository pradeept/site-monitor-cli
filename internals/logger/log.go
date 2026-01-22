package logger

import (
	l "log"
	"os"
)

func Logger() *l.Logger {
	log := l.New(os.Stdout, "", l.LstdFlags)
	log.SetFlags(l.LstdFlags | l.Lshortfile)
	return log
}
