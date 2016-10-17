// Package mapconv implements conversion to map[string]string of most types.
package mapconv

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToMap(value interface{}, prefix string) (m map[string]string, err error) {
	m = make(map[string]string)

	if value == nil {
		m[prefix] = "null"
		return
	}

	assignSubValue := func(path string, v reflect.Value) (err error) {
		switch v.Kind() {
		case reflect.Bool:
			m[path] = strconv.FormatBool(v.Bool())
		case reflect.Float64:
			m[path] = strconv.FormatFloat(v.Float(), 'f', 0, 64)
		case reflect.Int:
			m[path] = strconv.FormatInt(v.Int(), 10)
		case reflect.String:
			m[path] = v.String()
		case reflect.Map, reflect.Slice, reflect.Interface:
			var subMap map[string]string
			subMap, err = ToMap(v.Interface(), path)
			for key, value := range subMap {
				m[key] = value
			}
		default:
			err = fmt.Errorf("mapconv: unknown value on path (%s) of type (%s)", path, v.Kind())
		}
		return err
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Map:
		if len(prefix) > 0 {
			prefix += "."
		}
		for _, k := range rv.MapKeys() {
			path := prefix + k.String()
			v := rv.MapIndex(k)
			err = assignSubValue(path, v)
			if err != nil {
				//return
			}
		}
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			path := prefix + "[" + strconv.Itoa(i+1) + "]"
			v := rv.Index(i)
			err = assignSubValue(path, v)
			if err != nil {
				//return
			}
		}
	default:
		err = assignSubValue(prefix, rv)
	}

	return m, err
}
