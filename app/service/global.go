package service

import (
	"github.com/alioth-center/foreseen/app/dao"
	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/config"
	"github.com/alioth-center/infrastructure/database/postgres"
	"github.com/alioth-center/infrastructure/logger"
)

var (
	log logger.Logger

	templateDatabase *dao.TemplateDB

	syncModels = []any{
		&model.Account{}, &model.Client{}, &model.Integration{}, &model.Task{}, &model.Template{}, &model.User{},
	}
)

func init() {
	log = logger.Default()

	databaseConfig := postgres.Config{}
	loadErr := config.LoadConfigWithKeys(&databaseConfig, "./config/config.yaml", "database")
	if loadErr != nil {
		panic("failed to load database config: " + loadErr.Error())
	}

	db, initDatabaseErr := postgres.NewPostgresDb(databaseConfig, syncModels...)
	if initDatabaseErr != nil {
		panic("failed to initialize database: " + initDatabaseErr.Error())
	}

	gorm := db.ExtMethods().GetGorm()
	templateDatabase = dao.NewTemplateDB(gorm)
}
