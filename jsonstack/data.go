package helpers

type KeyValue struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type ArrayElement struct {
	data     []interface{}
	dataType byte
}

func (ae *ArrayElement) append(data interface{}) {
	ae.data = append(ae.data, data)
}

func (ae *ArrayElement) hasIncompleteData() bool {
	length := len(ae.data)

	if length == 0 {
		return false
	}

	lastElement, ok := ae.data[length-1].(KeyValue)
	if ok {
		return lastElement.Value == nil
	}

	return false
}

func (ae *ArrayElement) supplyIncompleteData(data interface{}) {
	length := len(ae.data) - 1

	lastElement, ok := ae.data[length].(KeyValue)
	if ok {
		updateKeyValue := KeyValue{
			Key:   lastElement.Key,
			Value: data,
		}

		ae.data[length] = updateKeyValue
	}
}
