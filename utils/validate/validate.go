package validate

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Vali(source interface{}) error {
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.Indirect(reflect.ValueOf(source))

	if sourceType.Kind() == reflect.Ptr {
		sourceType = sourceType.Elem()
	}

	if sourceValue.Kind() == reflect.Struct {
		for i := 0; i < sourceValue.NumField(); i++ {
			value := sourceValue.Field(i)
			typ := sourceType.Field(i)

			z := iszero(value)
			j := typ.Tag.Get("json")
			if j == "-" {
				continue
			}

			if r := typ.Tag.Get("required"); r == "true" {
				if z {
					return fmt.Errorf("missing %s", typ.Name)
				}
			}

			if value.Kind() == reflect.Slice ||
				(value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Slice) {
				slice := value
				slice = reflect.Indirect(slice)
				for i := 0; i < slice.Len(); i++ {
					element := slice.Index(i)
					if element.Kind() == reflect.Struct ||
						(element.Kind() == reflect.Ptr && element.Elem().Kind() == reflect.Struct) {
						err := Vali(element.Interface())
						if err != nil {
							return err
						}
					}
				}
			}

			if value.Kind() == reflect.Struct ||
				(value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct) {
				err := Vali(value.Interface())
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	return errors.New(fmt.Sprintf("validate data must be struct, not %v", sourceValue.Kind()))
}

func iszero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		flag := true
		for i := 0; i < v.Len(); i++ {
			flag = flag && iszero(v.Index(i))
		}
		return flag
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(time.Time{}) {
			return v.Interface().(time.Time).IsZero()
		}

		flag := true
		for i := 0; i < v.NumField(); i++ {
			flag = flag && iszero(v.Field(i))
		}

		return flag
	default:
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
}
