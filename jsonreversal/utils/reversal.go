package utils

import jstack "jsonstack"

func ReverseArray(urlResponse interface{}) interface{} {
	convertedArray, ok := urlResponse.([]interface{})
	if !ok {
		return convertedArray
	}

	allocatedArray := make([]interface{}, len(convertedArray))

	for i := len(convertedArray) - 1; i >= 0; i-- {
		allocatedArray[len(convertedArray)-1-i] = ReverseUrlResponse(convertedArray[i])
	}

	return allocatedArray
}

func ReverseObject(urlResponse jstack.KeyValue) jstack.KeyValue {
	urlResponse.Key = ReverseString(urlResponse.Key)
	urlResponse.Value = ReverseUrlResponse(urlResponse.Value)

	return urlResponse
}

func ReverseString(urlResponse interface{}) string {
	convertedString, ok := urlResponse.(string)
	if !ok {
		return convertedString
	}

	reversed := make([]byte, len(convertedString))

	for i := len(convertedString) - 1; i >= 0; i-- {
		reversed[i] = convertedString[(len(convertedString)-1)-i]
	}

	return string(reversed)
}

func ReverseUrlResponse(urlResponse interface{}) interface{} {
	convertedMapValue, ok := urlResponse.(jstack.KeyValue)
	if ok {
		reversed := ReverseObject(convertedMapValue)
		return reversed
	}

	convertedArrayValue, ok := urlResponse.([]interface{})
	if ok {
		return ReverseArray(convertedArrayValue)
	}

	convertedStringValue, ok := urlResponse.(string)
	if ok {
		return ReverseString(convertedStringValue)
	}

	return urlResponse
}