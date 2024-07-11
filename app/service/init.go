package service

import (
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/infrastructure/logger"
	"time"
)

var (
	logging   logger.Logger
	baseField logger.Fields
)

func init() {
	logging = logger.NewCustomLoggerWithOpts(logger.WithCustomWriterOpts(
		logger.NewTimeBasedRotationFileWriter(entity.GlobalConfig.LogDir, func(time time.Time) (filename string) {
			return time.Format("2006-01-02") + "_stdout.jsonl"
		}),
	))

	baseField = logger.NewFields().WithService("foreseen")
}
