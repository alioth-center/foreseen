package dao

import (
	"context"

	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/database"
)

type IntegrationDB struct {
	core database.DatabaseV2
}

func NewIntegrationDB(core database.DatabaseV2) *IntegrationDB {
	return &IntegrationDB{core: core}
}

func (db *IntegrationDB) GetIntegrationByName(ctx context.Context, name string) (integration *model.Integration, err error) {
	integration = new(model.Integration)
	return integration, db.core.GetDataBySingleCondition(ctx, integration, model.IntegrationCols.Name, name)
}

func (db *IntegrationDB) CreateIntegration(ctx context.Context, integration *model.Integration) (created bool, err error) {
	return db.core.CreateSingleDataIfNotExist(ctx, integration)
}

func (db *IntegrationDB) UpdateIntegrationByName(ctx context.Context, name string, updates *model.Integration) (err error) {
	return db.core.UpdateDataBySingleCondition(ctx, updates, model.IntegrationCols.Name, name)
}

func (db *IntegrationDB) DeleteIntegrationByName(ctx context.Context, name string) (deleted bool, err error) {
	sess := db.core.GetGormCore(ctx).Delete(&model.Integration{}, &model.Integration{Name: name})
	return sess.RowsAffected > 0, sess.Error
}
