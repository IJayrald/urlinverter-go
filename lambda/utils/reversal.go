package utils

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

func ReverseObject(urlResponse interface{}) map[string]interface{} {
	convertedMap, ok := urlResponse.(map[string]interface{})
	if !ok {
		return convertedMap
	}

	allocatedMap := make(map[string]interface{})

	for key, value := range convertedMap {
		reversedKey := ReverseString(key)
		allocatedMap[reversedKey] = ReverseUrlResponse(value)
	}

	return allocatedMap
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
	convertedStringValue, ok := urlResponse.(string)
	if ok {
		return ReverseString(convertedStringValue)
	}

	convertedMapValue, ok := urlResponse.(map[string]interface{})
	if ok {
		reversed := ReverseObject(convertedMapValue)
		return reversed
	}

	convertedArrayValue, ok := urlResponse.([]interface{})
	if ok {
		return ReverseArray(convertedArrayValue)
	}

	return urlResponse
}
