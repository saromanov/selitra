package app

import log "github.com/Sirupsen/logrus"

// LogrusHook  provides hook for logrus
type LogrusHook struct {
	Lvs []log.Level
}

func (l LogrusHook) Levels() []log.Level {
	return l.Lvs
}

func (sh LogrusHook) Fire(e *log.Entry) error {
	return nil
}
