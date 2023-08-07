package helpers

import (
	"encoding/json"
	"io"
)

/*
Iterates the valid JSON token, and format according to its structure

Note: All tokens are stored in a slice to preserve the keys order
*/
func ParseTraverse(d *json.Decoder, js *JsonStack) error {
	token, err := d.Token()
	if err == io.EOF {
		return err
	}

	delim, ok := token.(json.Delim)
	if ok && (delim == '[' || delim == '{') {
		js.Push(ArrayElement{dataType: byte(delim)})
		return ParseTraverse(d, js)
	}

	if ok && (delim == ']' || delim == '}') {
		js.MergePop()
		return ParseTraverse(d, js)
	}

	top := js.Top()
	if top.dataType == '[' {
		top.append(token)
		return ParseTraverse(d, js)
	}

	if top.dataType == '{' {
		isOpen := top.hasIncompleteData()
		if isOpen {
			top.supplyIncompleteData(token)
			return ParseTraverse(d, js)
		}
		top.append(KeyValue{
			Key:   token,
			Value: nil,
		})
		return ParseTraverse(d, js)
	}

	return ParseTraverse(d, js)
}
