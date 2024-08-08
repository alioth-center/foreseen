package router

import (
	"github.com/alioth-center/foreseen/app/api"
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/infrastructure/network/http"
)

var IntegrationApiRouterGroup = []http.EndPointInterface{
	http.NewEndPointBuilder[*entity.GetIntegrationRequest, *entity.GetIntegrationResponse]().
		SetRouter(http.NewRouter("/v1/integration/:integration_name")).
		SetAllowMethods(http.GET).
		SetNecessaryHeaders(http.HeaderAuthorization).
		SetNecessaryParams("integration_name").
		SetHandlerChain(api.IntegrationApi.GetIntegration()).
		Build(),
	http.NewEndPointBuilder[*entity.CreateIntegrationRequest, *entity.CreateIntegrationResponse]().
		SetRouter(http.NewRouter("/v1/integration")).
		SetAllowMethods(http.POST).
		SetNecessaryHeaders(http.HeaderContentType, http.HeaderAuthorization).
		SetHandlerChain(api.IntegrationApi.CreateIntegration()).
		Build(),
	http.NewEndPointBuilder[*entity.UpdateIntegrationRequest, *entity.UpdateIntegrationResponse]().
		SetRouter(http.NewRouter("/v1/integration/:integration_name")).
		SetAllowMethods(http.PUT).
		SetNecessaryHeaders(http.HeaderContentType, http.HeaderAuthorization).
		SetNecessaryParams("integration_name").
		SetHandlerChain(api.IntegrationApi.UpdateIntegration()).
		Build(),
	http.NewEndPointBuilder[*entity.DeleteIntegrationRequest, *entity.DeleteIntegrationResponse]().
		SetRouter(http.NewRouter("/v1/integration/:integration_name")).
		SetAllowMethods(http.DELETE).
		SetNecessaryHeaders(http.HeaderAuthorization).
		SetNecessaryParams("integration_name").
		SetHandlerChain(api.IntegrationApi.DeleteIntegration()).
		Build(),
}
