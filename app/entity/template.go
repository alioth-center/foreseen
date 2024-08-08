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

type UpdateTemplateRequest struct {
	Content   string          `json:"content" vc:"key:content,required"`
	Arguments json.RawMessage `json:"arguments,omitempty" vc:"key:arguments"`
}

type UpdateTemplateResult struct {
	TemplateID int             `json:"template_id"`
	Name       string          `json:"name"`
	Content    string          `json:"content"`
	Arguments  json.RawMessage `json:"arguments"`
	Preview    string          `json:"preview"`
}

type UpdateTemplateResponse = BaseResponse

type DeleteTemplateRequest = http.NoBody

type DeleteTemplateResult = http.NoResponse

type DeleteTemplateResponse = BaseResponse

type GetTemplatePreviewRequest = json.RawMessage

type GetTemplatePreviewResult struct {
	Preview string `json:"preview"`
}

type GetTemplatePreviewResponse = BaseResponse
