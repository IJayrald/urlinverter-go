package jsonreversal_test

import (
	"encoding/json"
	"strings"
	"testing"

	jreversal "jsonreversal"
	"jsonreversal/utils"
	jstack "jsonstack"
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
	expected := []interface{}{
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

// Tests array key-value reversal
func TestObjectReversal(t *testing.T) {
	testingData := []interface{}{
		jstack.KeyValue{Key: "name", Value: "Golang"},
		jstack.KeyValue{Key: "age", Value: 14},
		jstack.KeyValue{Key: "creator", Value: "Google Inc."},
	}

	expected := []jstack.KeyValue{
		{Key: "rotaerc", Value: ".cnI elgooG"},
		{Key: "ega", Value: 14},
		{Key: "eman", Value: "gnaloG"},
	}

	arbitraryValue := utils.ReverseUrlResponse(testingData)

	convertedInterface, ok := arbitraryValue.([]interface{})
	if !ok {
		t.Fatal("Expected array cannot be converted")
	}

	for index := 0; index < len(convertedInterface)-1; index++ {
		reversedKeyValue, ok := convertedInterface[index].(jstack.KeyValue)
		if !ok {
			t.Fatal("Expected KeyValy cannot be converted")
		}

		if (reversedKeyValue.Key != expected[index].Key) || (reversedKeyValue.Value != expected[index].Value) {
			t.Fatal(utils.ReversedNotExpected)
		}
	}
}

// Tests JSON object of strings reversal
func TestJsonArrayStringReversal(t *testing.T) {
	testingData := `["golang","reactjs"]`

	expected := `["sjtcaer","gnalog"]`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests JSON object reversal
func TestJsonObjectStringReversal(t *testing.T) {
	testingData := `{"name":"golang jr.","age":7,"creator":"Google Inc."}`

	expected := `{"rotaerc":".cnI elgooG","ega":7,"eman":".rj gnalog"}`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests JSON array of numbers - float and int reversal
func TestJsonArrayOfNumberStringReversal(t *testing.T) {
	testingData := `[1,2,3,9.4,23,9]`

	expected := `[9,23,9.4,3,2,1]`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests JSON array of booleans reversal
func TestJsonArrayOfBooleanStringReversal(t *testing.T) {
	testingData := `[true,false,false,true,true,false,false,true,true,true]`

	expected := `[true,true,true,false,false,true,true,false,false,true]`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests JSON array object without sub-objects reversal
func TestJsonArrayOfObjectLevel0StringReversal(t *testing.T) {
	testingData := `[{"name":"golang","age":27,"creator":"Google Inc."},{"name":"java","age":13,"creator":"James Gosling"}]`

	expected := `[{"rotaerc":"gnilsoG semaJ","ega":13,"eman":"avaj"},{"rotaerc":".cnI elgooG","ega":27,"eman":"gnalog"}]`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests JSON array object with 1 sub-object reversal
func TestJsonObjectLevel1StringReversal(t *testing.T) {
	testingData := `[{"name":{"first":"John","middle":"Ham","last":"Smith"},"age":37}]`

	expected := `[{"ega":37,"eman":{"tsal":"htimS","elddim":"maH","tsrif":"nhoJ"}}]`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests JSON array object with 2 sub-objects reversal
func TestJsonObjectLevel2StringReversal(t *testing.T) {
	testingData := `[{"name":{"first":"John","middle":"Ham","last":"Smith"},"age":37,"skills":[{"description":"Can do software development"},{"description":"Able to construct cloud infrastructure"}]}]`

	expected := `[{"slliks":[{"noitpircsed":"erutcurtsarfni duolc tcurtsnoc ot elbA"},{"noitpircsed":"tnempoleved erawtfos od naC"}],"ega":37,"eman":{"tsal":"htimS","elddim":"maH","tsrif":"nhoJ"}}]`

	testingJson := &jstack.JsonStack{}
	json.Unmarshal([]byte(testingData), testingJson)

	testingJson.ReverseJson(utils.ReverseUrlResponse)

	reversedJson, err := json.Marshal(testingJson)
	if err != nil {
		t.Fatal(err)
	}

	if string(reversedJson) != expected {
		t.Fatal(utils.ReversedNotExpected)
	}
}

// Tests if the function handles existing url
func TestExistingUrl(t *testing.T) {
	mockServer := utils.MockServer

	testingUrl := mockServer.URL

	_, err := jreversal.HandleInvertUrlResponse(testingUrl)
	if err != nil {
		t.Fatalf("Error running \"handleInvertUrlResponse\" function: %s", err)
	}
}

// Tests if the function handles empty url
func TestEmptyUrl(t *testing.T) {
	testingUrl := ""

	_, err := jreversal.HandleInvertUrlResponse(testingUrl)
	if (err != nil) && !strings.Contains(err.Error(), "url should not be empty") {
		t.Fatalf("handleInvertUrlResponse function did not handle url key checking")
	}
}
