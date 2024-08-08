package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alioth-center/foreseen/app/entity"
	"github.com/alioth-center/foreseen/app/model"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/http"
	"github.com/alioth-center/infrastructure/utils/values"
)

// CheckToken checks the authorization token provided in the request header.
// If the token is invalid, the request is aborted and an unauthorized status
// code is returned along with an appropriate error message.
//
// Parameters:
//
//	ctx (http.Context[req, *entity.BaseResponse]): The context of the HTTP request,
//	                                                which includes request and response data.
func CheckToken[req any](ctx http.Context[req, *entity.BaseResponse]) {
	token := strings.TrimPrefix(ctx.HeaderParams().GetString(http.HeaderAuthorization), "Bearer ")
	if token != entity.GlobalConfig.Token {
		ctx.Abort()
		ctx.SetStatusCode(http.StatusUnauthorized)
		ctx.SetResponse(&entity.BaseResponse{ErrorCode: entity.ErrorCodeAuthorizationError, ErrorMessage: entity.ErrorMessageAuthorizationError})
		return
	}
}

// SuccessMessage sets a success message in the response if there is no error
// and the response data is not nil. This function should be called after the
// main processing logic to ensure that the response is correctly populated.
//
// Parameters:
//
//	ctx (http.Context[req, *entity.BaseResponse]): The context of the HTTP request,
//	                                                which includes request and response data.
func SuccessMessage[req any](ctx http.Context[req, *entity.BaseResponse]) {
	ctx.Next()
	if ctx.Response().ErrorCode == 0 && ctx.Response().ErrorMessage == "" && ctx.Response().Data != nil {
		ctx.SetResponse(&entity.BaseResponse{
			ErrorMessage: "success",
			Data:         ctx.Response().Data,
		})
	}
}

// RenderTemplatePreview renders a template using the provided arguments and
// returns the rendered result as a string. The function merges the default
// template arguments with the provided arguments, with the latter overwriting
// the former if there are conflicts.
//
// Parameters:
//
//	ctx (context.Context): The context for logging and other operations.
//	template (*model.Template): The template to be rendered, including its content and default arguments.
//	args (json.RawMessage): The new arguments to be used for rendering the template.
//
// Returns:
//
//	result (string): The rendered template as a string.
func RenderTemplatePreview(ctx context.Context, template *model.Template, args json.RawMessage) string {
	// template content: `Hello, ${name}!`
	// template arguments: `{"name": "world"}`
	// render result: `Hello, world!`
	argsRaw, argsNew, argsAct := map[string]any{}, map[string]any{}, map[string]string{}

	// parse default arguments
	unmarshalRawErr := json.Unmarshal(values.UnsafeStringToBytes(template.Arguments), &argsRaw)
	if unmarshalRawErr != nil {
		log.Error(logger.NewFields(ctx).WithMessage("render template preview error: invalid raw arguments").WithData(unmarshalRawErr.Error()))
	}

	// parse new arguments
	unmarshalNewErr := json.Unmarshal(args, &argsNew)
	if unmarshalNewErr != nil {
		log.Error(logger.NewFields(ctx).WithMessage("render template preview error: invalid actual arguments").WithData(unmarshalNewErr.Error()))
	}

	// merge arguments, new arguments will overwrite raw arguments
	for k, v := range argsRaw {
		argsAct[k] = fmt.Sprintf("%v", v)
	}
	for k, v := range argsNew {
		argsAct[k] = fmt.Sprintf("%v", v)
	}

	// render template
	return values.NewStringTemplateWithMap(template.Content, argsAct).Parse()
}

// RenderTemplate renders a template using the provided arguments and returns
// the rendered result as a string. The function first retrieves the template
// from the database by its name and then calls RenderTemplatePreview to render
// the template with the provided arguments.
//
// Parameters:
//
//	ctx (context.Context): The context for logging and other operations.
//	templateName (string): The name of the template to be rendered.
//	args (json.RawMessage): The arguments to be used for rendering the template.
//
// Returns:
//
//	success (bool): Whether the template was successfully rendered.
//	result (string): The rendered template as a string.
func RenderTemplate(ctx context.Context, templateName string, args json.RawMessage) (success bool, result string) {
	template, err := templateDatabase.GetTemplateByName(ctx, templateName)
	if err != nil {
		log.Error(logger.NewFields(ctx).WithMessage("render template error: failed to get template").WithData(err))
		return false, ""
	} else if template.ID == 0 || template.Content == "" {
		log.Info(logger.NewFields(ctx).WithMessage("render template error: template not exist").WithData(templateName))
		return false, ""
	}

	return true, RenderTemplatePreview(ctx, template, args)
}
