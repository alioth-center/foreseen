package entity

import (
	"encoding/json"
	"github.com/alioth-center/infrastructure/network/http"
)

type GetTemplateRequest http.NoBody

type GetTemplateResult struct {
	TemplateID int             `json:"template_id"`
	Name       string          `json:"name"`
	Content    string          `json:"content"`
	Preview    string          `json:"preview"`
	Arguments  json.RawMessage `json:"arguments"`
}

type GetTemplateResponse = BaseResponse

type CreateTemplateRequest struct {
	Name      string          `json:"name" vc:"key:name,required"`
	Content   string          `json:"content" vc:"key:content,required"`
	Arguments json.RawMessage `json:"arguments" vc:"key:arguments"`
}

type CreateTemplateResult struct {
	TemplateID int    `json:"template_id"`
	Name       string `json:"name"`
	Preview    string `json:"preview"`
}

type CreateTemplateResponse = BaseResponse

type GetTemplatePreviewRequest = json.RawMessage

type GetTemplatePreviewResult struct {
	Preview string `json:"preview"`
}

type GetTemplatePreviewResponse = BaseResponse
