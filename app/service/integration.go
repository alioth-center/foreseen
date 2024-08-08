package service

import (
	"time"

	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/http"
)

type IntegrationService struct{}

func NewIntegrationService() *IntegrationService {
	return &IntegrationService{}
}

func (srv *IntegrationService) GetIntegration(ctx http.Context[*entity.GetIntegrationRequest, *entity.GetIntegrationResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start get integration").WithData(ctx.PathParams()))

	// query integration from database
	integration, err := integrationDatabase.GetIntegrationByName(ctx, ctx.PathParams().GetString("integration_name"))
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(GetIntegrationDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("get integration by name failed").WithData(err))
		return
	}
	if integration.ID == 0 {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetResponse(GetIntegrationNotExistError)
		log.Info(logger.NewFields(ctx).WithMessage("get integration not exist").WithData(ctx.PathParams()))
		return
	}

	// initialize response entity
	result := &entity.GetIntegrationResult{
		IntegrationID:   int(integration.ID),
		IntegrationName: integration.Name,
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.GetIntegrationResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("get integration success").WithData(result))
}

func (srv *IntegrationService) CreateIntegration(ctx http.Context[*entity.CreateIntegrationRequest, *entity.CreateIntegrationResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start create integration").WithData(ctx.Request()))

	// parse secrets
	if len(ctx.Request().Secrets) > 4 {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetResponse(CreateIntegrationTooMuchSecretsError)
		log.Info(logger.NewFields(ctx).WithMessage("create integration secrets error").WithData(ctx.Request().Secrets))
		return
	}
	actualSecrets := make([]string, 4)
	for i, secret := range ctx.Request().Secrets {
		actualSecrets[i] = secret
	}

	// create integration
	integration := &model.Integration{
		ID:        0,
		Name:      ctx.Request().Name,
		Secret1:   actualSecrets[0],
		Secret2:   actualSecrets[1],
		Secret3:   actualSecrets[2],
		Secret4:   actualSecrets[3],
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
	created, err := integrationDatabase.CreateIntegration(ctx, integration)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(CreateIntegrationDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("create integration failed").WithData(err))
		return
	}
	if !created {
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetResponse(CreateIntegrationConflictError)
		log.Warn(logger.NewFields(ctx).WithMessage("create integration conflict").WithData(ctx.Request()))
		return
	}

	// initialize response entity
	result := &entity.CreateIntegrationResult{
		IntegrationID:   int(integration.ID),
		IntegrationName: integration.Name,
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.CreateIntegrationResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("create integration success").WithData(result))
}

func (srv *IntegrationService) UpdateIntegration(ctx http.Context[*entity.UpdateIntegrationRequest, *entity.UpdateIntegrationResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start update integration").WithData(ctx.PathParams()))

	// parse secrets
	if len(ctx.Request().Secrets) > 4 {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetResponse(UpdateIntegrationTooMuchSecretsError)
		log.Info(logger.NewFields(ctx).WithMessage("update integration secrets error").WithData(ctx.Request().Secrets))
		return
	}

	// build update database entity
	integration := &model.Integration{
		Secret1: ctx.Request().Secrets[0],
		Secret2: ctx.Request().Secrets[1],
		Secret3: ctx.Request().Secrets[2],
		Secret4: ctx.Request().Secrets[3],
	}
	err := integrationDatabase.UpdateIntegrationByName(ctx, ctx.PathParams().GetString("integration_name"), integration)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(UpdateIntegrationDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("update integration failed").WithData(err))
		return
	}

	// query integration from database
	integration, err = integrationDatabase.GetIntegrationByName(ctx, ctx.PathParams().GetString("integration_name"))
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(GetIntegrationDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("get integration by name failed").WithData(err))
		return
	}

	// initialize response entity
	result := &entity.UpdateIntegrationResult{
		IntegrationID:   int(integration.ID),
		IntegrationName: integration.Name,
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.UpdateIntegrationResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("update integration success").WithData(result))
}

func (srv *IntegrationService) DeleteIntegration(ctx http.Context[*entity.DeleteIntegrationRequest, *entity.DeleteIntegrationResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start delete integration").WithData(ctx.PathParams()))

	// delete integration from database
	deleted, err := integrationDatabase.DeleteIntegrationByName(ctx, ctx.PathParams().GetString("integration_name"))
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(DeleteIntegrationDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("delete integration failed").WithData(err))
		return
	}
	if !deleted {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetResponse(DeleteIntegrationNotExistError)
		log.Info(logger.NewFields(ctx).WithMessage("delete integration not exist").WithData(ctx.PathParams()))
		return
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.DeleteIntegrationResponse{Data: &entity.DeleteIntegrationResult{}})
	log.Info(logger.NewFields(ctx).WithMessage("delete integration success").WithData(ctx.PathParams()))
}

var (
	GetIntegrationDatabaseError = &entity.GetIntegrationResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	GetIntegrationNotExistError = &entity.GetIntegrationResponse{
		ErrorCode:    entity.ErrorCodeGetIntegrationNotExist,
		ErrorMessage: entity.ErrorMessageGetIntegrationNotExist,
	}

	CreateIntegrationTooMuchSecretsError = &entity.CreateIntegrationResponse{
		ErrorCode:    entity.ErrorCodeCreateIntegrationTooMuchSecrets,
		ErrorMessage: entity.ErrorMessageCreateIntegrationTooMuchSecrets,
	}
	CreateIntegrationDatabaseError = &entity.CreateIntegrationResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	CreateIntegrationConflictError = &entity.CreateIntegrationResponse{
		ErrorCode:    entity.ErrorCodeCreateIntegrationConflict,
		ErrorMessage: entity.ErrorMessageCreateIntegrationConflict,
	}

	UpdateIntegrationTooMuchSecretsError = &entity.UpdateIntegrationResponse{
		ErrorCode:    entity.ErrorCodeUpdateIntegrationTooMuchSecrets,
		ErrorMessage: entity.ErrorMessageUpdateIntegrationTooMuchSecrets,
	}
	UpdateIntegrationDatabaseError = &entity.UpdateIntegrationResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}

	DeleteIntegrationDatabaseError = &entity.DeleteIntegrationResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	DeleteIntegrationNotExistError = &entity.DeleteIntegrationResponse{
		ErrorCode:    entity.ErrorCodeDeleteIntegrationNotExist,
		ErrorMessage: entity.ErrorMessageDeleteIntegrationNotExist,
	}
)
