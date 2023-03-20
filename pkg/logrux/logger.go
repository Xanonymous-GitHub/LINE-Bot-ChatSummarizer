package logrux

import (
	"github.com/kkdai/LINE-Bot-ChatSummarizer/pkg/env"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	if env.IsDebugMode {
		logger.SetLevel(logrus.DebugLevel)
	}

	logger.SetReportCaller(env.IsDebugMode)

	return logger
}
