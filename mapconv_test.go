package mapconv

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	tests := map[string]struct {
		prefix  string
		provide interface{}
		expect  map[string]string
	}{
		"nil": {
			provide: nil,
			expect: map[string]string{
				"": "null",
			},
		},
		"bool": {
			provide: true,
			expect: map[string]string{
				"": "true",
			},
		},
		"int": {
			provide: int(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"int8": {
			provide: int8(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"int16": {
			provide: int16(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"int32": {
			provide: int32(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"int64": {
			provide: int64(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"uint": {
			provide: uint(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"uint8": {
			provide: uint8(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"uint16": {
			provide: uint16(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"uint32": {
			provide: uint32(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"uint64": {
			provide: uint64(42),
			expect: map[string]string{
				"": "42",
			},
		},
		"float32": {
			provide: float32(4.2),
			expect: map[string]string{
				"": "4.2",
			},
		},
		"float64": {
			provide: float64(4.2),
			expect: map[string]string{
				"": "4.2",
			},
		},
		"string": {
			provide: "foo",
			expect: map[string]string{
				"": "foo",
			},
		},
		"map of maps of strings": {
			provide: map[string]map[string]string{
				"foo": map[string]string{
					"bar": "baz",
				},
				"qux": map[string]string{
					"quux":   "corge",
					"grault": "garply",
				},
			},
			expect: map[string]string{
				`["foo"]["bar"]`:    "baz",
				`["qux"]["quux"]`:   "corge",
				`["qux"]["grault"]`: "garply",
			},
		},
		"map of strings": {
			provide: map[string]string{
				"foo": "foo text",
				"bar": "bar text",
			},
			expect: map[string]string{
				`["foo"]`: "foo text",
				`["bar"]`: "bar text",
			},
		},
		"map of slices of strings": {
			provide: map[string][]string{
				"names":  []string{"bill", "bob"},
				"colors": []string{"blue", "green", "red"},
			},
			expect: map[string]string{
				`["names"][1]`:  "bill",
				`["names"][2]`:  "bob",
				`["colors"][1]`: "blue",
				`["colors"][2]`: "green",
				`["colors"][3]`: "red",
			},
		},
		"slice of ints": {
			provide: []int{9, 99, 999},
			expect: map[string]string{
				`[1]`: "9",
				`[2]`: "99",
				`[3]`: "999",
			},
		},
		"array of ints": {
			provide: [...]int{9, 99, 999},
			expect: map[string]string{
				`[1]`: "9",
				`[2]`: "99",
				`[3]`: "999",
			},
		},
		"slice of strings": {
			prefix:  "strings",
			provide: []string{"foo", "bar", "baz"},
			expect: map[string]string{
				`strings[1]`: "foo",
				`strings[2]`: "bar",
				`strings[3]`: "baz",
			},
		},
		"slice of maps": {
			provide: []map[string]string{
				{"name": "foo"},
				{"name": "bar"},
			},
			expect: map[string]string{
				`[1]["name"]`: "foo",
				`[2]["name"]`: "bar",
			},
		},
		"slice of slice of ints": {
			provide: [][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			expect: map[string]string{
				"[1][1]": "1",
				"[1][2]": "2",
				"[2][1]": "3",
				"[2][2]": "4",
			},
		},
	}
	jsonDoc := `{
		"name": "bob",
		"age": 35,
		"children": [
			{
				"name": "jack",
				"age": 5
			},
			{
				"name": "jill",
				"age": 7
			}
		]
	}`
	jsonExpect := map[string]string{
		`["name"]`:                "bob",
		`["age"]`:                 "35",
		`["children"][1]["name"]`: "jack",
		`["children"][1]["age"]`:  "5",
		`["children"][2]["name"]`: "jill",
		`["children"][2]["age"]`:  "7",
	}
	var jsonValue interface{}
	err := json.Unmarshal([]byte(jsonDoc), &jsonValue)
	if err != nil {
		t.Errorf("error unmarshalling test value: %v", err)
		t.FailNow()
	}
	tests["json.Unmarshalled value"] = struct {
		prefix  string
		provide interface{}
		expect  map[string]string
	}{
		provide: jsonValue,
		expect:  jsonExpect,
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := ToMap(test.provide, test.prefix)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("\nexpected: %#v \n but got: %#v", test.expect, actual)
			}
		})
	}
}
