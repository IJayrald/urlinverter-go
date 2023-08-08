package jsonstack

import (
	"bytes"
	"errors"
	"fmt"
)

func bundleObject(b *bytes.Buffer, data []interface{}) error {
	b.WriteRune('{')
	for index, sliceElement := range data {
		keyValue, ok := sliceElement.(KeyValue)
		if !ok {
			return errors.New("Cannot convert to KeyValue struct")
		}
		key, ok := keyValue.Key.(string)
		if !ok {
			return errors.New("Cannot convert \"Key\" to string")
		}
		bundleString(b, key)
		b.WriteRune(':')
		err := Bundle(b, keyValue.Value)
		if err != nil {
			return err
		}
		if index < len(data)-1 {
			b.WriteRune(',')
		}
	}
	b.WriteRune('}')
	return nil
}

func bundleArray(b *bytes.Buffer, data []interface{}) error {
	b.WriteRune('[')
	for index, sliceElement := range data {
		err := Bundle(b, sliceElement)
		if err != nil {
			return err
		}
		if index < len(data)-1 {
			b.WriteRune(',')
		}
	}
	b.WriteRune(']')
	return nil
}

func bundleString(b *bytes.Buffer, data string) error {
	b.WriteRune('"')
	b.WriteString(data)
	b.WriteRune('"')
	return nil
}

// Bundles the parsed JSON into []byte. Basically, reversing the parsed JSON into its original format
func Bundle(b *bytes.Buffer, data interface{}) error {
	slicedValue, ok := data.([]interface{})
	if ok {
		if len(slicedValue) == 0 {
			return bundleArray(b, []interface{}{})
		}
		_, ok := slicedValue[0].(KeyValue)
		if ok {
			return bundleObject(b, slicedValue)
		}
		return bundleArray(b, slicedValue)
	}

	textValue, ok := data.(string)
	if ok {
		return bundleString(b, textValue)
	}

	if data == nil {
		b.WriteString("null")
		return nil
	}

	b.WriteString(fmt.Sprintf("%v", data))

	return nil
}
