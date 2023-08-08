package utils

import "fmt"

type InvalidMethod string

func (s InvalidMethod) Error() string {
	return fmt.Sprintf("HTTP Invalid Method: %s", string(s))
}

type BadRequest string

func (b BadRequest) Error() string {
	return fmt.Sprintf("HTTP Bad Request: %s", string(b))
}
