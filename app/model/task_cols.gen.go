// Code generated by alioth-center/database-columns. DO NOT EDIT.
// Code generated by alioth-center/database-columns. DO NOT EDIT.
// Code generated by alioth-center/database-columns. DO NOT EDIT.

package model

type taskCols struct {
	ID            string
	IntegrationID string
	UserID        string
	ClientID      string
	AccountID     string
	TemplateID    string
	Argument      string
	Status        string
	CreatedAt     string
	UpdatedAt     string
}

var TaskCols = &taskCols{
	ID:            "id",
	IntegrationID: "integration_id",
	UserID:        "user_id",
	ClientID:      "client_id",
	AccountID:     "account_id",
	TemplateID:    "template_id",
	Argument:      "argument",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
