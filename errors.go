package mapconv

import (
	"fmt"
	"reflect"
)

type unsupportedKindError struct {
	path string
	kind reflect.Kind
}

func (err unsupportedKindError) Error() string {
	return fmt.Sprintf("mapconv: unsupported value on path (%s) of type (%s)", err.path, err.kind)
}

type unknownKindError struct {
	path string
	kind reflect.Kind
}

func (err unknownKindError) Error() string {
	return fmt.Sprintf("mapconv: unknown value on path (%s) of type (%s)", err.path, err.kind)
}
