package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"urlinverter.com/inverter/utils"
)

func getBody(r io.Reader) interface{} {
	requestBody, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var parsedBody interface{}

	json.Unmarshal(requestBody, &parsedBody)

	return parsedBody
}

func handleJsonResponse(url string) (io.Reader, error) {
	response, err := http.Get(url)
	if !strings.Contains(response.Header.Get("Content-Type"), "application/json") {
		return nil, utils.BadRequest("Not Valid JSON response")
	}
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func sendJsonResponse(jsonResponse []byte, statusCode int, w http.ResponseWriter) error {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}

func handleInvertUrlResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte(utils.HttpInvalidMethod))
		return
	}

	HttpBody := getBody(r.Body)
	urlValue := ""

	parsedBody, ok := HttpBody.(map[string]interface{})
	if ok {
		value, ok := parsedBody["url"]
		if !ok {
			w.Write([]byte(utils.HttpBadRequest))
			return
		}

		urlValue, ok = value.(string)
	} else {
		log.Printf("Invalid request body: %s", HttpBody)
		sendJsonResponse([]byte(utils.HttpBadRequest), http.StatusAccepted, w)
		return
	}

	jsonBody, err := handleJsonResponse(urlValue)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	HttpBody = getBody(jsonBody)
	reversed := utils.ReverseUrlResponse(HttpBody)

	response := utils.Inversion{
		Message: utils.Success,
		Details: utils.Details{
			Response: HttpBody,
			Inverted: reversed,
		},
	}

	jsonStringified, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling the JSON: %s", err)
		sendJsonResponse([]byte(utils.HttpInternalServerError), http.StatusInternalServerError, w)
		return
	}

	sendJsonResponse(jsonStringified, http.StatusAccepted, w)
}

func main() {
	http.HandleFunc("/invert", handleInvertUrlResponse)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
