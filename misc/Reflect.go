package misc

import (
	"reflect"
	"errors"
)

type Reflection map[string]interface{}

func Reflect(arg interface{}) Reflection {
	reflection := Reflection{}

	value := reflect.ValueOf(arg)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	for i := 0; i < value.NumField(); i++ {
		reflection[value.Type().Field(i).Name] = value.Field(i).Interface()
	}
	return reflection
}
func Clone(source interface{}, destin interface{}) error {
	destinValue := reflect.ValueOf(destin)
	if destinValue.Kind() != reflect.Ptr && destinValue.Kind() != reflect.Map {
		return errors.New("destin must be ptr")
	}

	sourceValue := reflect.ValueOf(source)
	if sourceValue.Kind() == reflect.Ptr {
		destinValue.Elem().Set(sourceValue.Elem())
	} else if sourceValue.Kind() == reflect.Map {
		for _, key := range sourceValue.MapKeys() {
			value := sourceValue.MapIndex(key)
			destinValue.SetMapIndex(key, value)
		}

	} else {
		destinValue.Elem().Set(sourceValue)
	}

	return nil
}