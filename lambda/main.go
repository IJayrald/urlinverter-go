package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
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
	if !strings.Contains(response.Header.Get(utils.ContentType), utils.Json) {
		return nil, utils.BadRequest(utils.JsonResponseNotValid)
	}
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func handleInvertUrlResponse(ctxt context.Context, lambdaInput map[string]interface{}) (utils.LambdaResponse, error) {
	url, ok := lambdaInput[utils.Url]
	if !ok {
		return utils.LambdaResponse{}, utils.BadRequest(utils.UrlNotExisting)
	}

	convertedUrl, ok := url.(string)
	if !ok {
		return utils.LambdaResponse{}, utils.BadRequest(utils.UrlNotValidText)
	}

	jsonBody, err := handleJsonResponse(convertedUrl)
	if err != nil {
		return utils.LambdaResponse{}, utils.BadRequest(err.Error())
	}

	HttpBody := getBody(jsonBody)
	reversed := utils.ReverseUrlResponse(HttpBody)

	return utils.LambdaResponse{
		Message: utils.Success,
		Details: utils.Details{
			Response: HttpBody,
			Inverted: reversed,
		},
	}, nil
}

func main() {
	lambda.Start(handleInvertUrlResponse)
}
