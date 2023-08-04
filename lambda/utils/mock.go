package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/jaswdr/faker"
)

var MockData = faker.New()

var MockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(ContentType, Json)

	serialized, err := json.Marshal(map[string]interface{}{
		"id":    MockData.UUID(),
		"name":  MockData.Person().Name(),
		"email": MockData.Person().Contact().Email,
		"address": map[string]interface{}{
			"houseNumber": MockData.Address().BuildingNumber(),
			"country":     MockData.Address().Country(),
			"city":        MockData.Address().City(),
		},
		"phoneNumber": MockData.Phone(),
	})
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(serialized)
}))
