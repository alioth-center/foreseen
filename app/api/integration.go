package api

import (
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/service"
	"github.com/alioth-center/infrastructure/network/http"
)

var IntegrationApi IntegrationApiImpl

type IntegrationApiImpl struct {
	srv *service.IntegrationService
}

func (impl IntegrationApiImpl) GetIntegration() http.Chain[*entity.GetIntegrationRequest, *entity.GetIntegrationResponse] {
	return http.NewChain[*entity.GetIntegrationRequest, *entity.GetIntegrationResponse](
		service.CheckToken[*entity.GetIntegrationRequest],
		service.SuccessMessage[*entity.GetIntegrationRequest],
		impl.srv.GetIntegration,
	)
}

func (impl IntegrationApiImpl) CreateIntegration() http.Chain[*entity.CreateIntegrationRequest, *entity.CreateIntegrationResponse] {
	return http.NewChain[*entity.CreateIntegrationRequest, *entity.CreateIntegrationResponse](
		service.CheckToken[*entity.CreateIntegrationRequest],
		service.SuccessMessage[*entity.CreateIntegrationRequest],
		impl.srv.CreateIntegration,
	)
}

func (impl IntegrationApiImpl) UpdateIntegration() http.Chain[*entity.UpdateIntegrationRequest, *entity.UpdateIntegrationResponse] {
	return http.NewChain[*entity.UpdateIntegrationRequest, *entity.UpdateIntegrationResponse](
		service.CheckToken[*entity.UpdateIntegrationRequest],
		service.SuccessMessage[*entity.UpdateIntegrationRequest],
		impl.srv.UpdateIntegration,
	)
}

func (impl IntegrationApiImpl) DeleteIntegration() http.Chain[*entity.DeleteIntegrationRequest, *entity.DeleteIntegrationResponse] {
	return http.NewChain[*entity.DeleteIntegrationRequest, *entity.DeleteIntegrationResponse](
		service.CheckToken[*entity.DeleteIntegrationRequest],
		service.SuccessMessage[*entity.DeleteIntegrationRequest],
		impl.srv.DeleteIntegration,
	)
}
