package apperrors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message  string
	Code     string
	HTTPCode int
}

var (
	EnvConfigLoadError = AppError{
		Message: "failed to load env file",
		Code:    "ENV_INIT_ERR",
	}

	EnvConfigVarError = AppError{
		Message: "config path hasn't been found in environment variables",
		Code:    "ENV_CONFIG_VAR_ERR",
	}

	EnvConfigParseError = AppError{
		Message: "failed to parse env file",
		Code:    "ENV_PARSE_ERR",
	}

	LoggerInitError = AppError{
		Message: "cannot init logger",
		Code:    "LOGGER_INIT_ERR",
	}

	InsertionFailedErr = AppError{
		Message:  "insertion operation has been failed",
		Code:     "INSERTION_ERR_FAILED",
		HTTPCode: http.StatusInternalServerError,
	}

	ClientConnectionGRPCServer = AppError{
		Message:  "client connection of grpc server has been failed",
		Code:     "GRPC_CONECTION_ERR_FAILED",
		HTTPCode: http.StatusInternalServerError,
	}

	GRPCErr = AppError{
		Message:  "can not get value from grpc server",
		Code:     "GETIING_DATA_FROM_GRPC_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	HashingErr = AppError{
		Message:  "Can not hash a password",
		Code:     "HASHING_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	CanNotCreateTokenErr = AppError{
		Message:  "can't create token",
		Code:     "TOKEN_CREATE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	WrongPasswordErr = AppError{
		Message:  "wrong password",
		Code:     "WRONG_PASSWORD_ERR",
		HTTPCode: http.StatusForbidden,
	}

	CanNotBindErr = AppError{
		Message:  "couldn't bind some data",
		Code:     "BINDING_ERR",
		HTTPCode: http.StatusBadRequest,
	}

	ExtractTokenErr = AppError{
		Message:  "can not extract token from header",
		Code:     "EXTRACT_TOKEN_ERR",
		HTTPCode: http.StatusUnauthorized,
	}

	ParseTokenErr = AppError{
		Message:  "can not parse token",
		Code:     "PARSE_TOKEN_ERR",
		HTTPCode: http.StatusUnauthorized,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message:  fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:     appError.Code,
		HTTPCode: appError.HTTPCode,
	}
}

func Is(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
