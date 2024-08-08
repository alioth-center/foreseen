package dao

import (
	"context"

	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/database"
)

type TemplateDB struct {
	core database.DatabaseV2
}

func NewTemplateDB(core database.DatabaseV2) *TemplateDB {
	return &TemplateDB{core: core}
}

func (db *TemplateDB) GetTemplateByName(ctx context.Context, name string) (template *model.Template, err error) {
	template = new(model.Template)
	return template, db.core.GetDataBySingleCondition(ctx, template, model.TemplateCols.Name, name)
}

func (db *TemplateDB) CreateTemplate(ctx context.Context, template *model.Template) (created bool, err error) {
	return db.core.CreateSingleDataIfNotExist(ctx, template)
}

func (db *TemplateDB) UpdateTemplateByName(ctx context.Context, name string, updates *model.Template) (err error) {
	return db.core.UpdateDataBySingleCondition(ctx, updates, model.TemplateCols.Name, name)
}

func (db *TemplateDB) DeleteTemplateByName(ctx context.Context, name string) (deleted bool, err error) {
	sess := db.core.GetGormCore(ctx).Delete(&model.Template{}, &model.Template{Name: name})
	return sess.RowsAffected > 0, sess.Error
}
