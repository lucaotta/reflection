package main

import (
	"reflect"
	"strconv"
)

func marshalInt(val reflect.Value) []byte {
	s := strconv.FormatInt(val.Int(), 10)
	return []byte(s)
}

func marshalUint(val reflect.Value) []byte {
	s := strconv.FormatUint(val.Uint(), 10)
	return []byte(s)
}

func Marshal(input interface{}) (result []byte, err error) {
	inputType := reflect.TypeOf(input)
	switch inputType.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		result = append(result, marshalInt(reflect.ValueOf(input))...)

	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		result = append(result, marshalUint(reflect.ValueOf(input))...)
	}

	return
}

func main() {

}
