package service

import (
	"encoding/json"
	"fmt"
	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/http"
	"github.com/alioth-center/infrastructure/utils/values"
	"time"
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
	}

	// build preview content
	argumentBuffer := values.UnsafeStringToBytes(template.Arguments)
	argumentsRaw, arguments := map[string]any{}, map[string]string{}
	decodeErr := json.Unmarshal(argumentBuffer, &argumentsRaw)
	if decodeErr != nil {
		// unmarshal arguments failed, report error
		log.Error(logger.NewFields(ctx).WithMessage("decode template arguments failed").WithData(decodeErr))
	} else {
		// convert arguments to string
		for key, value := range argumentsRaw {
			arguments[key] = fmt.Sprintf("%v", value)
		}
	}
	result.Arguments = argumentBuffer
	result.Preview = values.NewStringTemplateWithMap(template.Content, arguments).Parse()

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

	// build database entity
	template := &model.Template{
		Name:      ctx.Request().Name,
		Content:   ctx.Request().Content,
		Arguments: values.UnsafeBytesToString(ctx.Request().Arguments),
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
	}

	// build preview content
	argumentsRaw, arguments := map[string]any{}, map[string]string{}
	decodeErr := json.Unmarshal(ctx.Request().Arguments, &argumentsRaw)
	if decodeErr != nil {
		// unmarshal arguments failed, report error
		log.Error(logger.NewFields(ctx).WithMessage("decode template arguments failed").WithData(decodeErr))
	} else {
		// convert arguments to string
		for key, value := range argumentsRaw {
			arguments[key] = fmt.Sprintf("%v", value)
		}
	}
	result.Preview = values.NewStringTemplateWithMap(ctx.Request().Content, arguments).Parse()

	// set response
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetResponse(&entity.CreateTemplateResponse{Data: result})
	log.Info(logger.NewFields(ctx).WithMessage("create template success").WithData(result))
}

func (srv *TemplateService) GetTemplatePreview(ctx http.Context[*entity.GetTemplatePreviewRequest, *entity.GetTemplatePreviewResponse]) {
	log.Info(logger.NewFields(ctx).WithMessage("start get template preview").WithData(ctx.PathParams()))

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
	result := &entity.GetTemplatePreviewResult{}

	// build preview content
	var argumentBuffer = values.UnsafeStringToBytes(template.Arguments)
	if len(*ctx.Request()) > 0 {
		argumentBuffer = *ctx.Request()
	}
	argumentsRaw, arguments := map[string]any{}, map[string]string{}
	decodeErr := json.Unmarshal(argumentBuffer, &argumentsRaw)
	if decodeErr != nil {
		// unmarshal arguments failed, report error
		log.Error(logger.NewFields(ctx).WithMessage("decode template arguments failed").WithData(decodeErr))
	} else {
		// convert arguments to string
		for key, value := range argumentsRaw {
			arguments[key] = fmt.Sprintf("%v", value)
		}
	}
	result.Preview = values.NewStringTemplateWithMap(template.Content, arguments).Parse()

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

	CreateTemplateDatabaseError = &entity.CreateTemplateResponse{
		ErrorCode:    entity.ErrorCodeDatabaseError,
		ErrorMessage: entity.ErrorMessageDatabaseError,
	}
	CreateTemplateConflictError = &entity.CreateTemplateResponse{
		ErrorCode:    entity.ErrorCodeCreateTemplateConflict,
		ErrorMessage: entity.ErrorMessageCreateTemplateConflict,
	}
)
