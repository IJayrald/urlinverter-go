package utils

func ReverseArray(urlResponse interface{}) interface{} {
	value, ok := urlResponse.([]interface{})
	if !ok {
		return value
	}

	arr := make([]interface{}, len(value))

	for i := len(value) - 1; i >= 0; i-- {
		arr[len(value)-1-i] = ReverseUrlResponse(value[i])
	}

	return arr
}

func ReverseObject(urlResponse any) map[string]interface{} {
	value, ok := urlResponse.(map[string]interface{})
	if !ok {
		return value
	}

	arr := make(map[string]interface{})

	for key, keyVal := range value {
		reversedKey := ReverseString(key)
		arr[reversedKey] = ReverseUrlResponse(keyVal)
	}

	return arr
}

func ReverseString(urlResponse any) string {
	value, ok := urlResponse.(string)
	if !ok {
		return value
	}

	reversed := make([]byte, len(value))

	for i := len(value) - 1; i >= 0; i-- {
		reversed[i] = value[(len(value)-1)-i]
	}

	return string(reversed)
}

func ReverseUrlResponse(urlResponse any) any {
	valueS, ok := urlResponse.(string)
	if ok {
		return ReverseString(valueS)
	}

	valueM, ok := urlResponse.(map[string]interface{})
	if ok {
		reversed := ReverseObject(valueM)
		return reversed
	}

	valueA, ok := urlResponse.([]interface{})
	if ok {
		return ReverseArray(valueA)
	}

	return urlResponse
}
