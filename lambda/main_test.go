package main

import (
	"context"
	"strings"
	"testing"

	"urlinverter.com/inverter/utils"
)

// Checks the string reversal
func TestStringReversal(t *testing.T) {
	testingData := "test"
	expected := "tset"

	reversedString := utils.ReverseString(testingData)

	if len(reversedString) == 0 {
		t.Fatalf("Reversed string is empty: '%s'", reversedString)
	}

	if reversedString != expected {
		t.Fatalf("%s is not reversed", reversedString)
	}
}

// Checks the array and value reversal
func TestArrayReversal(t *testing.T) {
	testingData := []interface{}{
		"ABC",
		"DEF",
		"GHI",
	}
	expected := []string{
		"IHG",
		"FED",
		"CBA",
	}

	reversedArray := utils.ReverseArray(testingData)

	convertedArray, ok := reversedArray.([]interface{})
	if !ok {
		t.Fatal("Returned array is not a valid array of string")
	}

	if len(convertedArray) == 0 {
		t.Fatalf("Empty reversed array: %s", reversedArray)
	}

	for index := 0; index < len(convertedArray); index++ {
		if convertedArray[index] != expected[index] {
			t.Fatalf("Value %s is not reversed", convertedArray[index])
		}
	}
}

// Checks the map key and value reversal
func TestObjectReversal(t *testing.T) {
	mockAge := utils.MockData.Int8Between(13, 100)
	mockPhone := utils.MockData.Person().Contact().Phone

	testingData := map[string]interface{}{
		"age":         mockAge,
		"phoneNumber": mockPhone,
	}

	expected := map[string]interface{}{
		"ega":         mockAge,
		"rebmuNenohp": utils.ReverseString(mockPhone),
	}

	reversedObject := utils.ReverseObject(testingData)

	if len(reversedObject) == 0 {
		t.Fatalf("Empty reversed object: %s", reversedObject)
	}

	for key, value := range reversedObject {
		expectedValue, ok := expected[key]
		if !ok {
			t.Fatalf("Key %s is not reversed", key)
		}

		if expectedValue != value {
			t.Fatalf("Value %s is not equal to expected value. Not reversed", expectedValue)
		}
	}
}

// Checks if lambda function responds to event
func TestUrlExistingLambdaFunction(t *testing.T) {
	mockServer := utils.MockServer

	testingData := map[string]interface{}{
		"url": mockServer.URL,
	}

	_, err := handleInvertUrlResponse(context.TODO(), testingData)
	if err != nil {
		t.Fatalf("Error running \"handleInvertUrlResponse\" function: %s", err)
	}
}

// Checks if url key is not existing
func TestUrlNotExistingLambdaFunction(t *testing.T) {
	mockServer := utils.MockServer

	testingData := map[string]interface{}{
		"site": mockServer.URL,
	}

	_, err := handleInvertUrlResponse(context.TODO(), testingData)
	if (err != nil) && !strings.Contains(err.Error(), utils.UrlNotExisting) {
		t.Fatalf("handleInvertUrlResponse function did not handle url key checking")
	}
}
