package main

import (
	"testing"

	"urlinverter.com/inverter/utils"
)

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

func TestObjectReversal(t *testing.T) {
	testingData := map[string]interface{}{
		"age":         27,
		"phoneNumber": "09211234567",
	}

	expected := map[string]interface{}{
		"ega":         27,
		"rebmuNenohp": "76543211290",
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
