package model

type Task struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement"`
	IntegrationID int64  `gorm:"column:integration_id;type:integer;notnull;index:idx_integration_id"`
	UserID        int64  `gorm:"column:user_id;type:integer;notnull;index:idx_user_id"`
	ClientID      int64  `gorm:"column:client_id;type:integer;notnull;index:idx_client_id"`
	AccountID     int64  `gorm:"column:account_id;type:integer;notnull;index:idx_account_id"`
	TemplateID    int64  `gorm:"column:template_id;type:integer;notnull;index:idx_template_id"`
	Argument      string `gorm:"column:argument;type:text;notnull"`
	Status        int32  `gorm:"column:status;type:integer;notnull;default:0;index:idx_status"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime"`
}

func (t *Task) TableName() string {
	return TableNameTasks
}
