package utils

import (
	"errors"
	"os"
	"reflect"
	"strconv"
)

func ParseEnv(val interface{}) error {
	rv := reflect.Indirect(reflect.ValueOf(val))

	if !rv.CanSet() {
		return errors.New("value could not be set")
	}

	return loopAndAssign(rv, "env", func(key string) string {
		return os.Getenv(key)
	})
}

func loopAndAssign(rv reflect.Value, tag string, getEnv func(string) string) error {
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := t.Field(i)
		if envTag, ok := field.Tag.Lookup(tag); ok {
			envValue := getEnv(envTag)
			if envValue == "" && field.Tag.Get("validate") == "required" {
				return errors.New("required environment variable not set: " + envTag)
			} else if envValue == "" {
				return nil
			}
			switch field.Type.Kind() {
			case reflect.Int: 
				num, err := strconv.Atoi(envValue)
				if err != nil {
					return err
				}
				rv.Field(i).Set(reflect.ValueOf(num))
			case reflect.String:
				rv.Field(i).Set(reflect.ValueOf(envValue))
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(envValue)
				if err != nil {
					return err
				}
				rv.Field(i).Set(reflect.ValueOf(boolVal))
			default:
				return errors.New("unsupported type: " + field.Type.String())
			}
		}
	}
	return nil
}