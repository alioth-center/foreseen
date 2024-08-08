package router

import (
	"github.com/alioth-center/foreseen/app/api"
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/infrastructure/network/http"
)

var TemplateApiRouterGroup = []http.EndPointInterface{
	http.NewEndPointBuilder[*entity.GetTemplateRequest, *entity.GetTemplateResponse]().
		SetRouter(http.NewRouter("/v1/template/:template_name")).
		SetAllowMethods(http.GET).
		SetNecessaryHeaders(http.HeaderAuthorization).
		SetNecessaryParams("template_name").
		SetHandlerChain(api.TemplateApi.GetTemplate()).
		Build(),
	http.NewEndPointBuilder[*entity.CreateTemplateRequest, *entity.CreateTemplateResponse]().
		SetRouter(http.NewRouter("/v1/template")).
		SetAllowMethods(http.POST).
		SetNecessaryHeaders(http.HeaderContentType, http.HeaderAuthorization).
		SetHandlerChain(api.TemplateApi.CreateTemplate()).
		Build(),
	http.NewEndPointBuilder[*entity.UpdateTemplateRequest, *entity.UpdateTemplateResponse]().
		SetRouter(http.NewRouter("/v1/template/:template_name")).
		SetAllowMethods(http.PUT).
		SetNecessaryHeaders(http.HeaderContentType, http.HeaderAuthorization).
		SetNecessaryParams("template_name").
		SetHandlerChain(api.TemplateApi.UpdateTemplate()).
		Build(),
	http.NewEndPointBuilder[*entity.DeleteTemplateResult, *entity.DeleteTemplateResponse]().
		SetRouter(http.NewRouter("/v1/template/:template_name")).
		SetAllowMethods(http.DELETE).
		SetNecessaryHeaders(http.HeaderAuthorization).
		SetNecessaryParams("template_name").
		SetHandlerChain(api.TemplateApi.DeleteTemplate()).
		Build(),
	http.NewEndPointBuilder[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse]().
		SetRouter(http.NewRouter("/v1/template/:template_name/preview")).
		SetAllowMethods(http.GET).
		SetNecessaryHeaders(http.HeaderAuthorization).
		SetNecessaryParams("template_name").
		SetHandlerChain(api.TemplateApi.GetTemplatePreview()).
		Build(),
}
