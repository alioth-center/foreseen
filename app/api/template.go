package api

import (
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/service"
	"github.com/alioth-center/infrastructure/network/http"
)

var TemplateApi TemplateApiImpl

type TemplateApiImpl struct {
	srv *service.TemplateService
}

func (impl TemplateApiImpl) GetTemplate() http.Chain[*entity.GetTemplateRequest, *entity.GetTemplateResponse] {
	return http.NewChain[*entity.GetTemplateRequest, *entity.GetTemplateResponse](
		service.CheckToken[*entity.GetTemplateRequest],
		impl.srv.GetTemplate,
	)
}

func (impl TemplateApiImpl) CreateTemplate() http.Chain[*entity.CreateTemplateRequest, *entity.CreateTemplateResponse] {
	return http.NewChain[*entity.CreateTemplateRequest, *entity.CreateTemplateResponse](
		service.CheckToken[*entity.CreateTemplateRequest],
		impl.srv.CreateTemplate,
	)
}

func (impl TemplateApiImpl) GetTemplatePreview() http.Chain[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse] {
	return http.NewChain[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse](
		service.CheckToken[*entity.GetTemplatePreviewRequest],
		impl.srv.GetTemplatePreview,
	)
}
