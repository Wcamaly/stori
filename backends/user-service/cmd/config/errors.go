package config

import (
	"context"
	"encoding/json"
	"net/http"
	"stori/user-service/pkg/config/errors"
	"stori/user-service/pkg/config/logs"
)

var errToStatusCode = map[*errors.ErrorCode]int{
	// TODO
}

func WriteErr(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	logs.Error(ctx, "service error", logs.WithError(err))

	errB, err := mapErr(err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	statusCode := getStatusCode(err)
	w.WriteHeader(statusCode)
	w.Write(errB)
}

func mapErr(err error) ([]byte, error) {
	if err, ok := err.(*errors.Error); ok {
		return json.Marshal(map[string]interface{}{
			"code":    err.Code().String(),
			"message": err.Message(),
		})
	}

	return json.Marshal(map[string]interface{}{
		"message": err.Error(),
	})
}

func getStatusCode(err error) int {
	if err, ok := err.(*errors.Error); ok {
		if code, ok := errToStatusCode[err.Code()]; ok {
			return code
		}
	}

	switch err.(type) {
	case *json.InvalidUnmarshalError, *json.SyntaxError, *json.UnmarshalFieldError, *json.UnmarshalTypeError:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
