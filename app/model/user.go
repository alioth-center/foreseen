package model

type User struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Type      int32  `gorm:"column:type;type:integer;notnull;default:0;index:idx_type"`
	Name      string `gorm:"column:name;type:varchar(255);notnull;uniqueIndex:idx_name"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime"`
}

func (u *User) TableName() string {
	return TableNameUsers
}
