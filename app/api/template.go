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
		service.SuccessMessage[*entity.GetTemplateRequest],
		impl.srv.GetTemplate,
	)
}

func (impl TemplateApiImpl) CreateTemplate() http.Chain[*entity.CreateTemplateRequest, *entity.CreateTemplateResponse] {
	return http.NewChain[*entity.CreateTemplateRequest, *entity.CreateTemplateResponse](
		service.CheckToken[*entity.CreateTemplateRequest],
		service.SuccessMessage[*entity.CreateTemplateRequest],
		impl.srv.CreateTemplate,
	)
}

func (impl TemplateApiImpl) UpdateTemplate() http.Chain[*entity.UpdateTemplateRequest, *entity.UpdateTemplateResponse] {
	return http.NewChain[*entity.UpdateTemplateRequest, *entity.UpdateTemplateResponse](
		service.CheckToken[*entity.UpdateTemplateRequest],
		service.SuccessMessage[*entity.UpdateTemplateRequest],
		impl.srv.UpdateTemplate,
	)
}

func (impl TemplateApiImpl) DeleteTemplate() http.Chain[*entity.DeleteTemplateRequest, *entity.DeleteTemplateResponse] {
	return http.NewChain[*entity.DeleteTemplateRequest, *entity.DeleteTemplateResponse](
		service.CheckToken[*entity.DeleteTemplateRequest],
		service.SuccessMessage[*entity.DeleteTemplateRequest],
		impl.srv.DeleteTemplate,
	)
}

func (impl TemplateApiImpl) GetTemplatePreview() http.Chain[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse] {
	return http.NewChain[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse](
		service.CheckToken[*entity.GetTemplatePreviewRequest],
		service.SuccessMessage[*entity.GetTemplatePreviewRequest],
		impl.srv.GetTemplatePreview,
	)
}
