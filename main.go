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

func marshalArray(val reflect.Value) ([]byte, error) {
	result := []byte{'['}
	for i := 0; i < val.Len(); i++ {
		res, err := Marshal(val.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		result = append(result, res...)
		result = append(result, ',')
	}
	if val.Len() > 0 {
		result[len(result)-1] = ']'
	} else {
		result = append(result, ']')
	}

	return result, nil
}

func marshalStruct(val reflect.Value) ([]byte, error) {
	result := []byte{'{'}
	inputType := val.Type()
	exportedFields := 0
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		if field.PkgPath == "" {
			// This is an exported field
			tag := field.Tag.Get("json")
			fieldName := field.Name
			if tag != "" {
				fieldName = tag
			}
			result = append(result, []byte(fieldName)...)
			result = append(result, ':')
			res, err := Marshal(val.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			result = append(result, res...)
			result = append(result, ',')
			exportedFields += 1
		}
	}

	if exportedFields > 0 {
		result[len(result)-1] = '}'
	} else {
		result = append(result, '}')
	}

	return result, nil
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

	case reflect.Array:
		fallthrough
	case reflect.Slice:
		res, e := marshalArray(reflect.ValueOf(input))
		if e != nil {
			err = e
			return
		}
		result = append(result, res...)

	case reflect.Struct:
		res, e := marshalStruct(reflect.ValueOf(input))
		if e != nil {
			err = e
			return
		}
		result = append(result, res...)
	}

	return
}

func main() {
}
