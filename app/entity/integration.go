package entity

import "github.com/alioth-center/infrastructure/network/http"

type GetIntegrationRequest http.NoBody

type GetIntegrationResult struct {
	IntegrationID   int    `json:"integration_id"`
	IntegrationName string `json:"integration_name"`
}

type GetIntegrationResponse = BaseResponse

type CreateIntegrationRequest struct {
	Name    string   `json:"name" vc:"key:name,required"`
	Secrets []string `json:"secrets" vc:"key:secrets,required"`
}

type CreateIntegrationResult struct {
	IntegrationID   int    `json:"integration_id"`
	IntegrationName string `json:"integration_name"`
}

type CreateIntegrationResponse = BaseResponse

type UpdateIntegrationRequest struct {
	Secrets []string `json:"secrets" vc:"key:secrets,required"`
}

type UpdateIntegrationResult struct {
	IntegrationID   int    `json:"integration_id"`
	IntegrationName string `json:"integration_name"`
}

type UpdateIntegrationResponse = BaseResponse

type DeleteIntegrationRequest http.NoBody

type DeleteIntegrationResult http.NoResponse

type DeleteIntegrationResponse = BaseResponse
