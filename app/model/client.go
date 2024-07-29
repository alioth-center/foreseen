package model

type Client struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name;type:varchar(255);notnull;uniqueIndex:idx_name"`
	Secret    string `gorm:"column:secret;type:varchar(255);notnull;uniqueIndex:idx_secret"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Client) TableName() string {
	return TableNameClients
}
