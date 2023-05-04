package apperrors

import (
	"fmt"
)

type AppError struct {
	Message string
	Code    string
}

var (
	ConfigReadErr = AppError{
		Message: "couldn't read config",
		Code:    "CONFIG_READ_ERR",
	}

	DbOpenErr = AppError{
		Message: "can't open mariadb",
		Code:    "DATEBASE_OPEN_ERR",
	}

	DbPingErr = AppError{
		Message: "can't pass ping verification",
		Code:    "DATEBASE_PING_ERR",
	}

	MigrateDriverErr = AppError{
		Message: "can't create migration driver wiht instance",
		Code:    "MIGRATE_DRIVER_ERR",
	}

	MigrateCreationErr = AppError{
		Message: "can't create new migration wiht instance",
		Code:    "MIGRATE_CREATION_ERR",
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:    appError.Code,
	}
}
