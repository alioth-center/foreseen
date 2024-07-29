package model

type Integration struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;type:varchar(255);notnull;uniqueIndex:idx_name"`
	Secret1   string `gorm:"column:secret1;type:varchar(255)"`
	Secret2   string `gorm:"column:secret2;type:varchar(255)"`
	Secret3   string `gorm:"column:secret3;type:varchar(255)"`
	Secret4   string `gorm:"column:secret4;type:varchar(255)"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime"`
}

func (i *Integration) TableName() string {
	return TableNameIntegrations
}
