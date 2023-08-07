package utils

import (
	"encoding/json"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func (er *ErrorResponse) Error() string {
	errorResponse := ErrorResponse{
		Message: er.Message,
		Details: er.Details,
	}

	response, _ := json.Marshal(errorResponse)

	return string(response)
}
