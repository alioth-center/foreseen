package dao

import (
	"context"
	"github.com/alioth-center/foreseen/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TemplateDB struct {
	core *gorm.DB
}

func NewTemplateDB(core *gorm.DB) *TemplateDB {
	return &TemplateDB{core: core}
}

func (db *TemplateDB) GetTemplateByName(ctx context.Context, name string) (template *model.Template, err error) {
	template = &model.Template{}
	return template, db.core.WithContext(ctx).Model(&model.Template{}).Where(model.TemplateCols.Name, name).Scan(template).Error
}

func (db *TemplateDB) CreateTemplate(ctx context.Context, template *model.Template) (created bool, err error) {
	session := db.core.WithContext(ctx).Model(&model.Template{}).Clauses(clause.OnConflict{DoNothing: true}).Create(template)
	return session.RowsAffected > 0, session.Error
}
