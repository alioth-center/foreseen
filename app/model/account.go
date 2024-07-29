package model

type Account struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement;uniqueIndex:idx_account_ids"`
	Integration int64  `gorm:"column:integration;type:integer;notnull;index:idx_integration;uniqueIndex:idx_account_ids"`
	UserID      int64  `gorm:"column:user_id;type:integer;notnull;index:idx_user_id;uniqueIndex:idx_account_ids"`
	Account     string `gorm:"column:account;type:varchar(255);notnull;index:idx_account"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime"`
}

func (a *Account) TableName() string {
	return TableNameAccounts
}
