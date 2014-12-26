package main

import "fmt"

type Processor struct{}

var processor = Processor{}

type JsonObject map[string]interface{}
type JsonCollection []JsonObject
type JsonArray []interface{}

func (p *Processor) Get(in interface{}, key string) (interface{}, error) {
	obj, err := p.toObject(in)

	if err != nil {
		return nil, err
	}

	return obj[key], nil
}

func (p *Processor) Select(in interface{}, keys ...string) (JsonCollection, error) {
	coll, err := p.toCollection(in)

	if err != nil {
		return nil, err
	}

	out := make(JsonCollection, 0)

	for _, obj := range coll {
		newobj := p.selectProp(obj, keys...)

		if len(newobj) != 0 {
			out = append(out, newobj)
		}
	}

	return out, nil
}

func (p *Processor) Reject(in interface{}, keys ...string) (JsonCollection, error) {
	coll, err := p.toCollection(in)

	if err != nil {
		return nil, err
	}

	out := make(JsonCollection, 0)

	for _, obj := range coll {
		newobj := p.rejectProp(obj, keys...)

		if len(newobj) != 0 {
			out = append(out, newobj)
		}
	}

	return out, nil
}

func (p *Processor) Head(in interface{}, length int) (JsonArray, error) {
	arr, err := p.toArray(in)

	if err != nil {
		return nil, err
	}

	return p.sliceArray(arr, 0, length), nil
}

func (p *Processor) Tail(in interface{}, length int) (JsonArray, error) {
	arr, err := p.toArray(in)

	if err != nil {
		return nil, err
	}

	return p.sliceArray(arr, len(arr)-length, len(arr)), nil
}

func (p *Processor) selectProp(obj JsonObject, keys ...string) JsonObject {
	ret := make(JsonObject)

	for key, val := range obj {
		for _, k := range keys {
			if key == k {
				ret[key] = val
			}
		}
	}

	return ret
}

func (p *Processor) rejectProp(obj JsonObject, keys ...string) JsonObject {
	ret := make(JsonObject)

	for key, val := range obj {
		hasKey := false
		for _, k := range keys {
			if key == k {
				hasKey = true
				break
			}
		}
		if hasKey == false {
			ret[key] = val
		}
	}

	return ret
}

func (p *Processor) sliceArray(arr JsonArray, start, length int) JsonArray {
	maxLen := len(arr)

	if start < 0 {
		start = 0
	}

	if start > maxLen {
		start = maxLen
	}

	end := start + length

	if end > maxLen {
		end = maxLen
	}

	if end < 0 {
		end = 0
	}

	return arr[start:end]
}

func (p *Processor) toObject(in interface{}) (JsonObject, error) {
	_obj, ok := in.(map[string]interface{})

	if !ok {
		return nil, p.invalidTypeError("Object")
	}

	var obj JsonObject = _obj

	return obj, nil
}

func (p *Processor) toArray(in interface{}) (JsonArray, error) {
	_arr, ok := in.([]interface{})

	if !ok {
		return nil, p.invalidTypeError("Array")
	}

	var arr JsonArray = _arr

	return arr, nil
}

func (p *Processor) toCollection(in interface{}) (JsonCollection, error) {
	arr, ok := in.([]interface{})

	if !ok {
		return nil, p.invalidTypeError("Collection")
	}

	coll := make(JsonCollection, len(arr))

	for i, val := range arr {
		obj, ok := val.(map[string]interface{})

		if !ok {
			return nil, p.invalidTypeError("Collection")
		}

		coll[i] = obj
	}

	return coll, nil
}

func (p *Processor) invalidTypeError(expected string) error {
	return fmt.Errorf("Invalid type error: expected type is %s", expected)
}
