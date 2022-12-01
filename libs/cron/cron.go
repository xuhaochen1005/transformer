package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"transformer_mz/api/design"
	"transformer_mz/libs/log"
)

type CLog struct {
	clog *logrus.Logger
}

func (l *CLog) Info(msg string, keysAndValues ...interface{}) {
	l.clog.WithFields(logrus.Fields{
		"data": keysAndValues,
	}).Info(msg)
}

func (l *CLog) Error(err error, msg string, keysAndValues ...interface{}) {
	l.clog.WithFields(logrus.Fields{
		"msg":  msg,
		"data": keysAndValues,
	}).Warn(msg)
}

func InitCron() {
	logger := &CLog{clog: log.Logger}
	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(logger)))
	c.AddFunc("@every 10s", func() {
		design.CheckDesignResults()
	})
	c.Start()
}
