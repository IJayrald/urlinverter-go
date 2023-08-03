package utils

type Details struct {
	Response any `json:"response"`
	Inverted any `json:"inverted"`
}

type Inversion struct {
	Message string  `json:"message"`
	Details Details `json:"details"`
}
