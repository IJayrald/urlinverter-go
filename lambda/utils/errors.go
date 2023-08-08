package utils

import (
	"encoding/json"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func (er *ErrorResponse) Error() string {
	response, _ := json.Marshal(er)

	return string(response)
}
