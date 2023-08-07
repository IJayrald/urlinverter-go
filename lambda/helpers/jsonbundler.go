package helpers

import (
	"bytes"
	"errors"
	"fmt"
)

func bundleObject(b bytes.Buffer, data []KeyValue) ([]byte, error) {
	b.WriteRune('{')
	for _, keyValue := range data {
		key, ok := keyValue.Key.(string)
		if !ok {
			return nil, errors.New("Cannot convert Key to string")
		}
		b.WriteString(key)
		b.WriteRune(':')
		converted, err := Bundle(b, keyValue)
		if err != nil {
			return nil, err
		}
		b.Write(converted)
		b.WriteRune(',')
	}
	b.WriteRune('}')
	b.WriteRune(',')
	return b.Bytes(), nil
}

func bundleArray(b bytes.Buffer, data []interface{}) ([]byte, error) {
	b.WriteRune('[')
	for _, value := range data {
		buffered, err := Bundle(b, value)
		if err != nil {
			return nil, err
		}
		b.Write(buffered)
		b.WriteRune(',')
	}
	b.WriteRune(']')
	fmt.Println(b.String())
	return b.Bytes(), nil
}

// Bundles the parsed JSON into []byte. Basically, reversing the parsed JSON into its original format
func Bundle(b bytes.Buffer, data interface{}) ([]byte, error) {
	slicedKeyValue, ok := data.([]KeyValue)
	if ok {
		return bundleObject(b, slicedKeyValue)
	}

	slicedInterface, ok := data.([]interface{})
	if ok {
		return bundleArray(b, slicedInterface)
	}

	textValue, ok := data.(string)
	if ok {
		b.WriteRune('"')
		b.WriteString(textValue)
		b.WriteRune('"')
	}

	return b.Bytes(), nil
}
