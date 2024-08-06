package router

import (
	"encoding/json"
	"github.com/alioth-center/foreseen/app/api"
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/infrastructure/network/http"
)

func init() {
	engine := http.NewEngine("/foreseen")

	engine.AddEndPoints(
		http.NewEndPointBuilder[*entity.LarkNotifyRequest, *entity.BaseResponse]().
			SetRouter(http.NewRouter("/v1/notify/lark")).
			SetAllowMethods(http.POST).
			SetNecessaryHeaders(http.HeaderContentType, http.HeaderAuthorization).
			SetHandlerChain(api.NotifyApi.NotifyLark()).
			Build(),
	)

	engine.AddEndPoints(
		http.NewEndPointBuilder[any, json.RawMessage]().
			SetRouter(http.NewRouter("/v1/callback")).
			SetAllowMethods(http.POST, http.GET).
			SetHandlerChain(api.CallbackApi.Callback()).
			Build(),
	)

	engine.AddEndPoints(
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
		http.NewEndPointBuilder[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse]().
			SetRouter(http.NewRouter("/v1/template/:template_name/preview")).
			SetAllowMethods(http.GET).
			SetNecessaryHeaders(http.HeaderAuthorization).
			SetNecessaryParams("template_name").
			SetHandlerChain(api.TemplateApi.GetTemplatePreview()).
			Build(),
	)

	engine.ServeAsync("0.0.0.0:8881", make(chan struct{}, 1))
}
