package main

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, actual, expected interface{}, msg string) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%s\nactual: %#v\nexpected: %#v", msg, actual, expected)
	}
}

func TestProcessorGet(t *testing.T) {
	in := map[string]interface{}{"foo": "bar"}

	out, err := processor.Get(in, "foo")

	if err != nil {
		t.Fatal(err)
	}

	assert(t, out, "bar", "get")
}

func TestProcessorSelect(t *testing.T) {
	in := []interface{}{
		map[string]interface{}{"key1": "1-1", "key2": "2-1", "key3": "3-1"},
		map[string]interface{}{"key1": "1-2", "key2": "2-2", "key3": "3-2"},
		map[string]interface{}{"key1": "1-3", "key2": "2-3"},
	}

	out, err := processor.Select(in, "key1", "key3")

	if err != nil {
		t.Fatal(err)
	}

	expected := JsonCollection{
		JsonObject{"key1": "1-1", "key3": "3-1"},
		JsonObject{"key1": "1-2", "key3": "3-2"},
		JsonObject{"key1": "1-3"},
	}

	assert(t, out, expected, "select")
}

func TestProcessorReject(t *testing.T) {
	in := []interface{}{
		map[string]interface{}{"key1": "1-1", "key2": "2-1", "key3": "3-1"},
		map[string]interface{}{"key1": "1-2", "key2": "2-2", "key3": "3-2"},
		map[string]interface{}{"key1": "1-3", "key2": "2-3"},
	}

	out, err := processor.Reject(in, "key1", "key3")

	if err != nil {
		t.Fatal(err)
	}

	expected := JsonCollection{
		JsonObject{"key2": "2-1"},
		JsonObject{"key2": "2-2"},
		JsonObject{"key2": "2-3"},
	}

	assert(t, out, expected, "reject")
}

func TestProcessorSlice(t *testing.T) {
	test := func(in []interface{}, start, length int, expected JsonArray, msg string) {
		out, err := processor.Slice(in, start, length)
		if err != nil {
			t.Fatal(err)
		}
		assert(t, out, expected, msg)
	}

	test([]interface{}{1, 2, 3, 4, 5}, 0, 2, JsonArray{1, 2}, "slice-1")
	test([]interface{}{1, 2, 3, 4, 5}, 2, 2, JsonArray{3, 4}, "slice-2")
	test([]interface{}{1, 2, 3, 4, 5}, 3, 10, JsonArray{4, 5}, "slice-3")
	test([]interface{}{1, 2, 3, 4, 5}, 3, 0, JsonArray{}, "slice-3")
	test([]interface{}{1, 2, 3, 4, 5}, 10, 10, JsonArray{}, "slice-4")
}

func TestProcessorhead(t *testing.T) {
	test := func(in []interface{}, length int, expected JsonArray, msg string) {
		out, err := processor.Head(in, length)
		if err != nil {
			t.Fatal(err)
		}
		assert(t, out, expected, msg)
	}

	test([]interface{}{1, 2, 3, 4, 5}, 0, JsonArray{}, "head-1")
	test([]interface{}{1, 2, 3, 4, 5}, 2, JsonArray{1, 2}, "head-2")
	test([]interface{}{1, 2, 3, 4, 5}, 5, JsonArray{1, 2, 3, 4, 5}, "head-3")
	test([]interface{}{1, 2, 3, 4, 5}, 10, JsonArray{1, 2, 3, 4, 5}, "head-4")
}

func TestProcessorTail(t *testing.T) {
	test := func(in []interface{}, length int, expected JsonArray, msg string) {
		out, err := processor.Tail(in, length)
		if err != nil {
			t.Fatal(err)
		}
		assert(t, out, expected, msg)
	}

	test([]interface{}{1, 2, 3, 4, 5}, 0, JsonArray{}, "tail-4")
	test([]interface{}{1, 2, 3, 4, 5}, 2, JsonArray{4, 5}, "tail-1")
	test([]interface{}{1, 2, 3, 4, 5}, 5, JsonArray{1, 2, 3, 4, 5}, "tail-2")
	test([]interface{}{1, 2, 3, 4, 5}, 10, JsonArray{1, 2, 3, 4, 5}, "tail-3")
}
