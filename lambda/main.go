package main

import (
	"context"
	jreversal "jsonreversal"
	"lambda/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEventRequest(ctx context.Context, lambdaInput map[string]interface{}) (interface{}, error) {
	arbitraryValue, ok := lambdaInput["url"]
	if !ok {
		return nil, &utils.ErrorResponse{
			Message: "Cannot find url property",
			Details: "Please pass a valid url",
		}
	}

	url, ok := arbitraryValue.(string)
	if !ok {
		return nil, &utils.ErrorResponse{
			Message: "Invalid url type",
			Details: "Please pass url in text type",
		}
	}

	return jreversal.HandleInvertUrlResponse(url)
}

func main() {
	lambda.Start(HandleLambdaEventRequest)
}
