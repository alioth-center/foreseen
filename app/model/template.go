package model

type Template struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;type:varchar(255);notnull;unique;uniqueIndex:idx_name"`
	Content   string `gorm:"column:content;type:text;notnull"`
	Arguments string `gorm:"column:arguments;type:text;notnull"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime"`
}

func (t *Template) TableName() string {
	return TableNameTemplates
}
