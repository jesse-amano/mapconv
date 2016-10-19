// Package mapconv implements conversion to map[string]string of most types.
package mapconv

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ToMap returns a map[path]value given any arbitrary value.
func ToMap(value interface{}, prefix string) (m map[string]string, err error) {
	m = make(map[string]string)

	if value == nil {
		m[prefix] = "null"
		return
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Map:
		for _, k := range rv.MapKeys() {
			path := prefix + `["` + k.String() + `"]`
			err = assignSubValue(m, path, rv.MapIndex(k).Interface())
			if err != nil {
				//return
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			path := prefix + "[" + strconv.Itoa(i+1) + "]"
			err = assignSubValue(m, path, rv.Index(i).Interface())
			if err != nil {
				//return
			}
		}
	default:
		err = assignSubValue(m, prefix, rv.Interface())
	}

	return m, err
}

func assignSubValue(m map[string]string, path string, v interface{}) (err error) {
	if v, ok := v.(fmt.Stringer); ok {
		m[path] = v.String()
		return
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Bool:
		m[path] = strconv.FormatBool(rv.Bool())
	case reflect.Float32, reflect.Float64:
		m[path] = strings.TrimRight(strings.TrimRight(strconv.FormatFloat(rv.Float(), 'f', 6, 64), "0"), ".")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		m[path] = strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		m[path] = strconv.FormatUint(rv.Uint(), 10)
	case reflect.String:
		m[path] = rv.String()
	case reflect.Array, reflect.Map, reflect.Slice:
		var subMap map[string]string
		subMap, err = ToMap(rv.Interface(), path)
		for key, value := range subMap {
			m[key] = value
		}
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Complex64, reflect.Complex128, reflect.Interface:
		err = unsupportedKindError{
			path: path,
			kind: rv.Kind(),
		}
	default:
		err = unknownKindError{
			path: path,
			kind: rv.Kind(),
		}
	}
	return err
}
