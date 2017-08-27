package util

import mylogger"github.com/gocomb/job-center/server/util/logger"

var Logger mylogger.NewLog

func init() {
	Logger.LogRegister(mylogger.LevelDebug)
}
