package utils

type Details struct {
	Original interface{} `json:"original"`
	Reversed interface{} `json:"reversed"`
}

type Response struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
