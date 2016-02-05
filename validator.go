package vision

import (
	"fmt"
	"reflect"
	"strings"
)

// Validate to validate provide parameters with a struct
func Validate(obj interface{}) (bool, error) {
	return ValidateCustom(tagName, obj)
}

// ValidateCustom to validate provide parameters with a struct
func ValidateCustom(tag string, obj interface{}) (bool, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false, fmt.Errorf("function only accepts structs; got %s", val.Kind())
	}
	for i := 0; i < val.NumField(); i++ {
		// value := val.Field(i)
		typef := val.Type().Field(i)
		tag := typef.Tag.Get(tag)
		if tag == "" || tag == "-" {
			continue
		}
		opts := extractOptions(tag)
		if len(opts) == 0 {
			continue
		}
		rules := extractRules(opts)
		if len(rules) == 0 {
			continue
		}
		fmt.Println("val:", rules)
	}
	return true, nil
}

func extractOptions(str string) []string {
	return strings.FieldsFunc(str, func(r rune) bool {
		switch r {
		case '|', ';', ',':
			return true
		}
		return false
	})
}

func extractRules(rules []string) map[string]string {
	opts := map[string]string{}
	for _, val := range rules {
		args := strings.SplitN(val, ":", -1)
		count := len(args)
		if count == 2 {
			opts[args[0]] = args[1]
		} else if count > 2 {
			panic("validation option(s) format are invalid")
		} else {
			opts[val] = "-"
		}
	}
	return opts
}
