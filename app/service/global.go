package service

import (
	"github.com/alioth-center/foreseen/app/dao"
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/config"
	"github.com/alioth-center/infrastructure/database/postgres"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/utils/values"
	"time"
)

var (
	log logger.Logger

	templateDatabase    *dao.TemplateDB
	integrationDatabase *dao.IntegrationDB

	syncModels = []any{
		&model.Account{}, &model.Client{}, &model.Integration{}, &model.Task{}, &model.Template{}, &model.User{},
	}
)

func init() {
	log = logger.NewCustomLoggerWithOpts(
		logger.WithCustomWriterOpts(
			logger.NewTimeBasedRotationFileWriter(entity.GlobalConfig.LogDir, func(time time.Time) (filename string) {
				return values.BuildStrings(time.Format("2006-01-02"), "_stdout.jsonl")
			}),
		),
	)

	databaseConfig := postgres.Config{}
	loadErr := config.LoadConfigWithKeys(&databaseConfig, "./config/config.yaml", "database")
	if loadErr != nil {
		panic("failed to load database config: " + loadErr.Error())
	}

	db, initDatabaseErr := postgres.NewPostgresSQLv2(databaseConfig, syncModels...)
	if initDatabaseErr != nil {
		panic("failed to initialize database: " + initDatabaseErr.Error())
	}

	templateDatabase = dao.NewTemplateDB(db)
	integrationDatabase = dao.NewIntegrationDB(db)
}
