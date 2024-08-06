package entity

const (
	ErrorCodeAuthorizationError = 4011
	ErrorCodeDatabaseError      = 5011

	ErrorCodeGetTemplateNotExist    = 4041
	ErrorCodeCreateTemplateConflict = 4091
)

const (
	ErrorMessageAuthorizationError = "invalid authorization"
	ErrorMessageDatabaseError      = "database error"

	ErrorMessageGetTemplateNotExist    = "template not exist"
	ErrorMessageCreateTemplateConflict = "template already exist"
)
