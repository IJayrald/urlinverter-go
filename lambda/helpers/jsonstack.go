package helpers

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

func (s *JsonStack) GetParsedJson() []interface{} {
	return s.Top().data
}

// Pushes the ArrayElement on the stack
func (s *JsonStack) Push(data ArrayElement) {
	s.data = append(s.data, data)
}

// Gets the topmost ArrayElement in the stack and returns it as a pointer
func (s *JsonStack) Top() *ArrayElement {
	return &s.data[len(s.data)-1]
}

// Removes the topmost ArrayElement in the stack. Commonly used when formatting process has finished
func (s *JsonStack) Pop() interface{} {
	var temp []interface{}

	index := len(s.data) - 1

	if index == 0 {
		return nil
	}

	temp = s.data[index].data

	s.data = s.data[:index]

	return temp
}

// Pops the stack and conditionally merge it to the previous ArrayElement
func (s *JsonStack) MergePop() {
	pop := s.Pop()
	top := s.Top()

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
func (s *JsonStack) UnmarshalJSON(b []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(b))

	err := ParseTraverse(decoder, s)

	return err
}

func (s *JsonStack) MarshalJSON() ([]byte, error) {
	var bufferedJson bytes.Buffer

	bundled, err := Bundle(bufferedJson, s.Top().data)

	return bundled, err
}

/*
Reverses the JSON order

Array => The elements order will be reversed

Object => The keys order will be reversed

String => The texts will be reversed
*/
func (s *JsonStack) ReverseJson(reversalFunc func(interface{}) interface{}) {
	reversed, ok := reversalFunc(s.GetParsedJson()).([]interface{})
	if ok {
		s.Top().data = reversed
	}
}
