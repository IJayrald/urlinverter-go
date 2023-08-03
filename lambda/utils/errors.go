package utils

type InvalidMethod string

func (s InvalidMethod) Error() string {
	return "HTTP Error [InvalidMethod]"
}

type BadRequest string

func (b BadRequest) Error() string {
	return "Bad Request"
}
