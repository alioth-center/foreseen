package service

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/http"
	"github.com/alioth-center/infrastructure/utils/values"
)

type TemplateService struct{}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

// GetTemplate get foreseen template info by template_name, see [api-ref] for more details
//
// [api-ref]: https://docs.alioth.center/foreseen-api-templates.html#-jl7xn3_19
func (srv *TemplateService) GetTemplate(ctx http.Context[*entity.GetTemplateRequest, *entity.GetTemplateResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start get template").WithData(ctx.PathParams()))

	// query template from database
	template, err := templateDatabase.GetTemplateByName(ctx, ctx.PathParams().GetString("template_name"))
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(GetTemplateDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("get template by name failed").WithData(err))
		return
	}
	if template.ID == 0 {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetResponse(GetTemplateNotExistError)
		log.Info(logger.NewFields(ctx).WithMessage("get template not exist").WithData(ctx.PathParams()))
		return
	}

	// initialize response entity
	result := &entity.GetTemplateResult{
		TemplateID: int(template.ID),
		Name:       template.Name,
		Content:    template.Content,
		Arguments:  values.UnsafeStringToBytes(template.Arguments),
		Preview:    RenderTemplatePreview(ctx, template, values.UnsafeStringToBytes(template.Arguments)),
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.GetTemplateResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("get template success").WithData(result))
}

// CreateTemplate create a new foreseen template, see [api-ref] for more details
//
// [api-ref]: https://docs.alioth.center/foreseen-api-templates.html#-jl7xn3_41
func (srv *TemplateService) CreateTemplate(ctx http.Context[*entity.CreateTemplateRequest, *entity.CreateTemplateResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start create template").WithData(ctx.Request()))

	// compact json arguments
	args := &bytes.Buffer{}
	compactErr := json.Compact(args, ctx.Request().Arguments)
	if compactErr != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetResponse(CreateTemplateArgumentsError)
		log.Info(logger.NewFields(ctx).WithMessage("create template arguments error").WithData(ctx.Request().Arguments))
	}

	// build database entity
	template := &model.Template{
		Name:      ctx.Request().Name,
		Content:   ctx.Request().Content,
		Arguments: values.UnsafeBytesToString(args.Bytes()),
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}

	// insert template to database
	created, err := templateDatabase.CreateTemplate(ctx, template)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(CreateTemplateDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("create template failed").WithData(err))
		return
	}
	if !created {
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetResponse(CreateTemplateConflictError)
		log.Warn(logger.NewFields(ctx).WithMessage("create template conflict").WithData(ctx.Request()))
		return
	}

	// initialize response entity
	result := &entity.CreateTemplateResult{
		TemplateID: int(template.ID),
		Name:       template.Name,
		Preview:    RenderTemplatePreview(ctx, template, ctx.Request().Arguments),
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.CreateTemplateResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("create template success").WithData(result))
}

// UpdateTemplate update foreseen template by template_name, see [api-ref] for more details
//
// [api-ref]: https://docs.alioth.center/foreseen-api-templates#-jl7xn3_65
func (srv *TemplateService) UpdateTemplate(ctx http.Context[*entity.UpdateTemplateRequest, *entity.UpdateTemplateResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start update template").WithData(ctx.PathParams()))

	// compact json arguments
	args := &bytes.Buffer{}
	compactErr := json.Compact(args, ctx.Request().Arguments)
	if compactErr != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetResponse(CreateTemplateArgumentsError)
		log.Info(logger.NewFields(ctx).WithMessage("create template arguments error").WithData(ctx.Request().Arguments))
	}

	// build update database entity
	template := &model.Template{
		Content:   ctx.Request().Content,
		Arguments: values.UnsafeBytesToString(args.Bytes()),
		UpdatedAt: time.Now().UnixMilli(),
	}

	// update template in database
	err := templateDatabase.UpdateTemplateByName(ctx, ctx.PathParams().GetString("template_name"), template)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(UpdateTemplateDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("update template failed").WithData(err))
		return
	}

	// query template from database
	template, err = templateDatabase.GetTemplateByName(ctx, ctx.PathParams().GetString("template_name"))
	if err != nil {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetResponse(UpdateTemplateNotExistError)
		log.Error(logger.NewFields(ctx).WithMessage("update template not exist").WithData(err))
		return
	}

	// build preview content
	result := &entity.UpdateTemplateResult{
		TemplateID: int(template.ID),
		Name:       template.Name,
		Content:    template.Content,
		Arguments:  ctx.Request().Arguments,
		Preview:    RenderTemplatePreview(ctx, template, ctx.Request().Arguments),
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.UpdateTemplateResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("update template success").WithData(ctx.PathParams()))
}

// DeleteTemplate delete foreseen template by template_name, see [api-ref] for more details
//
// [api-ref]: https://docs.alioth.center/foreseen-api-templates#-jl7xn3_90
func (srv *TemplateService) DeleteTemplate(ctx http.Context[*entity.DeleteTemplateRequest, *entity.DeleteTemplateResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start delete template").WithData(ctx.PathParams()))

	// delete template from database
	deleted, err := templateDatabase.DeleteTemplateByName(ctx, ctx.PathParams().GetString("template_name"))
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetResponse(DeleteTemplateDatabaseError)
		log.Error(logger.NewFields(ctx).WithMessage("delete template by name failed").WithData(err))
		return
	}
	if !deleted {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetResponse(DeleteTemplateNotExistError)
		log.Info(logger.NewFields(ctx).WithMessage("delete template not exist").WithData(ctx.PathParams()))
		return
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.DeleteTemplateResponse{Data: &entity.DeleteTemplateResult{}})
	log.Info(logger.NewFields(ctx).WithMessage("delete template success").WithData(ctx.PathParams()))
}

// GetTemplatePreview get foreseen template preview by template_name, see [api-ref] for more details
//
// [api-ref]: https://docs.alioth.center/foreseen-api-templates#-jl7xn3_113
func (srv *TemplateService) GetTemplatePreview(ctx http.Context[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start get template preview").WithData(ctx.PathParams()))

	// render template preview
	rendered, result := RenderTemplate(ctx, ctx.PathParams().GetString("template_name"), *ctx.Request())
	if !rendered {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetResponse(GetTemplatePreviewNotExistError)
		log.Info(logger.NewFields(ctx).WithMessage("get template preview not exist").WithData(ctx.PathParams()))
		return
	}

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.GetTemplatePreviewResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("get template preview success").WithData(result))
}

var (
	GetTemplateDatabaseError = &entity.GetTemplateResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	GetTemplateNotExistError = &entity.GetTemplateResponse{
		ErrorCode:    entity.ErrorCodeGetTemplateNotExist,
		ErrorMessage: entity.ErrorMessageGetTemplateNotExist,
	}

	CreateTemplateArgumentsError = &entity.CreateTemplateResponse{
		ErrorCode:    entity.ErrorCodeCreateTemplateArgumentsError,
		ErrorMessage: entity.ErrorMessageCreateTemplateArgumentsError,
	}
	CreateTemplateDatabaseError = &entity.CreateTemplateResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	CreateTemplateConflictError = &entity.CreateTemplateResponse{
		ErrorCode:    entity.ErrorCodeCreateTemplateConflict,
		ErrorMessage: entity.ErrorMessageCreateTemplateConflict,
	}

	UpdateTemplateDatabaseError = &entity.UpdateTemplateResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	UpdateTemplateNotExistError = &entity.UpdateTemplateResponse{
		ErrorCode:    entity.ErrorCodeUpdateTemplateNotExist,
		ErrorMessage: entity.ErrorMessageUpdateTemplateNotExist,
	}

	DeleteTemplateDatabaseError = &entity.DeleteTemplateResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	DeleteTemplateNotExistError = &entity.DeleteTemplateResponse{
		ErrorCode:    entity.ErrorCodeDeleteTemplateNotExist,
		ErrorMessage: entity.ErrorMessageDeleteTemplateNotExist,
	}

	GetTemplatePreviewNotExistError = &entity.GetTemplatePreviewResponse{
		ErrorCode:    entity.ErrorCodeGetTemplatePreviewNotExist,
		ErrorMessage: entity.ErrorMessageGetTemplatePreviewNotExist,
	}
)
