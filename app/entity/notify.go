package entity

type BaseResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Data         any    `json:"data,omitempty"`
}

type LarkNotifyRequest struct {
	Level          string   `json:"level"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	UserReceivers  []string `json:"user_receivers,omitempty"`
	GroupReceivers []string `json:"group_receivers,omitempty"`
}

type LarkNotifyResponse struct {
	Status string   `json:"status"`
	Errors []string `json:"errors,omitempty"`
}
