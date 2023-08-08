package jsonstack

import (
	"bytes"
	"encoding/json"
)

/*
Struct which handles the process of formatting the structure

The process pushes on the stack to maintain the structure
*/
type JsonStack struct {
	data []ArrayElement
}

func (js *JsonStack) GetParsedJson() interface{} {
	return js.Top().data
}

// Pushes the ArrayElement on the stack
func (js *JsonStack) Push(data ArrayElement) {
	js.data = append(js.data, data)
}

// Gets the topmost ArrayElement in the stack and returns it as a pointer
func (js *JsonStack) Top() *ArrayElement {
	return &js.data[len(js.data)-1]
}

// Removes the topmost ArrayElement in the stack. Commonly used when formatting process has finished
func (js *JsonStack) Pop() interface{} {
	var temp []interface{}

	index := len(js.data) - 1

	if index == 0 {
		return nil
	}

	temp = js.data[index].data

	js.data = js.data[:index]

	return temp
}

// Pops the stack and conditionally merge it to the previous ArrayElement
func (js *JsonStack) MergePop() {
	pop := js.Pop()
	top := js.Top()

	if pop == nil {
		return
	}

	isOpen := top.hasIncompleteData()
	if isOpen {
		top.supplyIncompleteData(pop)
		return
	}

	top.append(pop)
}

// json.Unmarshall calls this method when unmarshalled to the JsonStack struct
func (js *JsonStack) UnmarshalJSON(b []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(b))

	err := ParseTraverse(decoder, js)

	return err
}

func (js *JsonStack) MarshalJSON() ([]byte, error) {
	var bufferedJson bytes.Buffer

	err := Bundle(&bufferedJson, js.Top().data)

	return bufferedJson.Bytes(), err
}

/*
Reverses the JSON order

Array => The elements order will be reversed

Object => The keys order will be reversed

String => The texts will be reversed
*/
func (js *JsonStack) ReverseJson(reversalFunc func(interface{}) interface{}) {
	reversed, ok := reversalFunc(js.GetParsedJson()).([]interface{})
	if ok {
		js.Top().data = reversed
	}
}
