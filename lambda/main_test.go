package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mainlambda/utils"
	"strconv"
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

	convertedResponse, ok := response.(utils.Response)
	if !ok {
		t.Fatal(GeneralUnexpectedError, "Cannot convert to Response type")
	}

	buffered, err := json.Marshal(convertedResponse.Details)

	var details utils.Details

	json.Unmarshal(buffered, &details)

	var b bytes.Buffer

	unquotedOriginal, err := strconv.Unquote(fmt.Sprintf(`%q`, details.Original))
	unquotedReversed, err := strconv.Unquote(fmt.Sprintf(`%q`, details.Reversed))

	json.Indent(&b, []byte(unquotedOriginal), "", "\t")

	fmt.Println(b.String())

	b.Reset()
	json.Indent(&b, []byte(unquotedReversed), "", "\t")

	fmt.Println(b.String())

	if err != nil {
		t.Fatal(GeneralUnexpectedError, err)
	}
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
