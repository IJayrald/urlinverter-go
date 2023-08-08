package main

import (
	"context"
	jreversal "jsonreversal"
	"mainlambda/utils"

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

	if len(url) == 0 {
		return nil, &utils.ErrorResponse{
			Message: "Url is empty",
			Details: "Please pass url value",
		}
	}

	response, err := jreversal.HandleInvertUrlResponse(url)
	if err != nil {
		return nil, &utils.ErrorResponse{
			Message: "Error inverting URL response",
			Details: err.Error(),
		}
	}

	return utils.Response{
		Message: "Success inverting URL response",
		Details: response,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEventRequest)
}
