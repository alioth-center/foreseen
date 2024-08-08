package entity

const (
	ErrorCodeAuthorizationError = 4011
	ErrorCodeDatabaseError      = 5011

	ErrorCodeGetTemplateNotExist          = 4041
	ErrorCodeCreateTemplateArgumentsError = 4001
	ErrorCodeCreateTemplateConflict       = 4091
	ErrorCodeUpdateTemplateNotExist       = 4041
	ErrorCodeDeleteTemplateNotExist       = 4041
	ErrorCodeGetTemplatePreviewNotExist   = 4041

	ErrorCodeGetIntegrationNotExist          = 4041
	ErrorCodeCreateIntegrationTooMuchSecrets = 4001
	ErrorCodeCreateIntegrationConflict       = 4091
	ErrorCodeUpdateIntegrationTooMuchSecrets = 4001
	ErrorCodeDeleteIntegrationNotExist       = 4041
)

const (
	ErrorMessageAuthorizationError = "invalid authorization"
	ErrorMessageDatabaseError      = "database error"

	ErrorMessageGetTemplateNotExist          = "template not exist"
	ErrorMessageCreateTemplateArgumentsError = "invalid arguments"
	ErrorMessageCreateTemplateConflict       = "template already exist"
	ErrorMessageUpdateTemplateNotExist       = "template not exist"
	ErrorMessageDeleteTemplateNotExist       = "template not exist"
	ErrorMessageGetTemplatePreviewNotExist   = "template not exist"

	ErrorMessageGetIntegrationNotExist          = "integration not exist"
	ErrorMessageCreateIntegrationTooMuchSecrets = "too much secrets"
	ErrorMessageCreateIntegrationConflict       = "integration already exist"
	ErrorMessageUpdateIntegrationTooMuchSecrets = "too much secrets"
	ErrorMessageDeleteIntegrationNotExist       = "integration not exist"
)
