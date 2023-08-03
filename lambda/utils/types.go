package utils

type Details struct {
	Response any `json:"response"`
	Inverted any `json:"inverted"`
}

type LambdaResponse struct {
	Message string  `json:"message"`
	Details Details `json:"details"`
}
