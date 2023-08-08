package main

import (
	"context"
	"fmt"
	"mainlambda/utils"
	"strings"
	"testing"
)

const GeneralUnexpectedError = "Unexpected error"

func TestUrlPropertyExistsLambdaFunction(t *testing.T) {
	mockServer := utils.MockServer

	testingData := map[string]interface{}{
		"url": mockServer.URL,
	}

	response, err := HandleLambdaEventRequest(context.TODO(), testingData)
	if err != nil {
		t.Fatal(GeneralUnexpectedError, err)
	}
	fmt.Println(response)
}

func TestUrlPropertyNotExistsLambdaFunction(t *testing.T) {
	mockServer := utils.MockServer

	testingData := map[string]interface{}{
		"site": mockServer.URL,
	}

	_, err := HandleLambdaEventRequest(context.TODO(), testingData)
	if (err != nil) && !strings.Contains(err.Error(), "Cannot find url property") {
		t.Fatal(GeneralUnexpectedError, err)
	}
}

func TestUrlValueNumberTypeLambdaFunction(t *testing.T) {
	testingData := map[string]interface{}{
		"url": 10,
	}

	_, err := HandleLambdaEventRequest(context.TODO(), testingData)
	if (err != nil) && !strings.Contains(err.Error(), "Invalid url type") {
		t.Fatal(GeneralUnexpectedError, err)
	}
}

func TestUrlValueEmptyLambdaFunction(t *testing.T) {
	testingData := map[string]interface{}{
		"url": "",
	}

	_, err := HandleLambdaEventRequest(context.TODO(), testingData)
	if (err != nil) && !strings.Contains(err.Error(), "Url is empty") {
		t.Fatal(GeneralUnexpectedError, err)
	}
}
