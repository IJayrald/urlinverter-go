package utils

// HTTP-based Error
const (
	HttpInvalidMethod       = "Invalid Method Request"
	HttpBadRequest          = "Invalid Request"
	HttpInternalServerError = "Internal Server Error"
)

// HTTP-based Headers
const (
	ContentType = "Content-Type"
)

// Mime-Types
const (
	Json = "application/json"
)

// HTTP Request Key/Properties
const (
	Url = "url"
)

// Error Responses
const (
	UrlNotExisting       = "url key is not existing"
	UrlNotValidText      = "url is not valid text"
	JsonResponseNotValid = "Not Valid JSON response"
)

// Success Responses
const (
	Success = "Request Successful"
)
