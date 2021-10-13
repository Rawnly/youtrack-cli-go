package util

import (
	"fmt"
	"reflect"
)

func GetFieldsOf(t reflect.Type) string {
	fields := ""

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		if len(tag) == 0 {
			continue
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			value := GetFieldsOf(field.Type)

			if i == 0 {
				fields = fmt.Sprintf("%s(%s)", tag, value)
			} else {
				fields += fmt.Sprintf(",%s(%s)", tag, value)
			}

			continue
		case reflect.Slice:
			value := GetFieldsOf(field.Type.Elem())

			if i == 0 {
				fields = fmt.Sprintf("%s(%s)", tag, value)
			} else {
				fields += fmt.Sprintf(",%s(%s)", tag, value)
			}
		default:
			if len(fields) == 0 {
				fields = tag
				continue
			}

			fields += fmt.Sprintf(",%s", tag)
			break
		}
	}

	return fields
}
