package jsonreversal

import (
	"encoding/json"
	"io"
	"jsonreversal/utils"
	jstack "jsonstack"
	"log"
	"net/http"
	"strings"
)

func getBody(r io.Reader) []byte {
	requestBody, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return requestBody
}

func handleJsonResponse(url string) (io.Reader, error) {
	response, err := http.Get(url)
	if !strings.Contains(response.Header.Get(utils.ContentType), utils.Json) {
		return nil, utils.BadRequest(utils.JsonResponseNotValid)
	}
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func HandleInvertUrlResponse(url string) (utils.Response, error) {
	if len(strings.TrimSpace(url)) == 0 {
		return utils.Response{}, utils.BadRequest("url should not be empty")
	}

	jsonBody, err := handleJsonResponse(url)
	if err != nil {
		return utils.Response{}, utils.BadRequest(err.Error())
	}

	HttpBody := getBody(jsonBody)

	jsonStack := &jstack.JsonStack{}
	json.Unmarshal(HttpBody, &jsonStack)

	reversed, err := json.Marshal(jsonStack)
	if err != nil {
		return utils.Response{}, err
	}

	return utils.Response{
		Original: string(HttpBody),
		Reversed: string(reversed),
	}, nil
}
