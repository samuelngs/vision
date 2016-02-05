package vision

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Validate to validate provide parameters with a struct
func Validate(obj interface{}) (*Param, Error) {
	return ValidateCustom(tagName, obj)
}

// ValidateCustom to validate provide parameters with a struct
func ValidateCustom(tag string, obj interface{}) (*Param, Error) {
	param := &Param{m: make(map[string]interface{})}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, ValueError{fmt.Sprintf("function only accepts structs; got %s", val.Kind())}
	}
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
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
		rules.RLock()
		defer rules.RUnlock()
		if len(rules.m) == 0 {
			continue
		}
		str := fmt.Sprintf("%v", value)
		res, err := validateField(str, rules)
		if err != nil {
			rules.RLock()
			defer rules.RUnlock()
			if rules.m["error"] != "" {
				return nil, ValueError{rules.m["error"]}
			}
			return nil, ValueError{fmt.Sprintf("%v %v", typef.Name, err.Error())}
		}
		param.Set(typef.Name, res)
	}
	return param, nil
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

func extractRules(rules []string) *Rules {
	opts := &Rules{m: make(map[string]string)}
	for _, val := range rules {
		args := strings.SplitN(val, ":", -1)
		count := len(args)
		if count > 2 {
			panic("validation option(s) format are invalid")
		} else if count == 2 {
			opts.RLock()
			defer opts.RUnlock()
			opts.m[args[0]] = args[1]
		} else {
			opts.RLock()
			defer opts.RUnlock()
			opts.m[val] = "-"
		}
	}
	return opts
}

func validateField(s string, rules *Rules) (interface{}, Error) {
	format := ""
	val, err := ParseString(s)
	if err != nil {
		return nil, err
	}
	for rule, opt := range rules.m {
		switch rule {
		case "string", "text":
			format = "string"
		case "int", "integer", "number":
			format = "integer"
		case "float", "float32", "float64":
			format = "float"
		case "bool", "boolean":
			format = "bool"
		case "required":
			if val.IsEmpty() {
				return nil, ErrEmpty
			}
		case "strict-required":
			if val.MustEmpty() {
				return nil, ErrEmpty
			}
		case "email":
			if !val.IsEmail() {
				return nil, ErrInvalid
			}
		case "url":
			if !val.IsURL() {
				return nil, ErrInvalid
			}
		case "latitude":
			if !val.IsLatitude() {
				return nil, ErrInvalid
			}
		case "longitude":
			if !val.IsLongitude() {
				return nil, ErrInvalid
			}
		case "us-phonenumber":
			if !val.IsUSPhoneNumber() {
				return nil, ErrInvalid
			}
		case "uuid":
			if !val.IsUUID() {
				return nil, ErrInvalid
			}
		case "uuid3":
			if !val.IsUUID3() {
				return nil, ErrInvalid
			}
		case "uuid4":
			if !val.IsUUID4() {
				return nil, ErrInvalid
			}
		case "uuid5":
			if !val.IsUUID5() {
				return nil, ErrInvalid
			}
		case "base64":
			if !val.IsBase64() {
				return nil, ErrInvalid
			}
		case "datauri":
			if !val.IsDataURI() {
				return nil, ErrInvalid
			}
		case "hexcolor":
			if !val.IsHexColor() {
				return nil, ErrInvalid
			}
		case "rgbcolor":
			if !val.IsRGBColor() {
				return nil, ErrInvalid
			}
		case "creditcard":
			if !val.IsCreditCard() {
				return nil, ErrInvalid
			}
		case "numberic":
			if !val.IsNumberic() {
				return nil, ErrInvalid
			}
		case "min":
			if opt == "" {
				panic("min requires a number")
			}
			i, err := strconv.ParseInt(fmt.Sprintf("%v", opt), 10, 64)
			if err != nil {
				panic(err)
			}
			if !val.IsLength(int(i), 0) {
				return nil, ErrMin
			}
		case "max":
			if opt == "" {
				panic("max requires a number")
			}
			i, err := strconv.ParseInt(fmt.Sprintf("%v", opt), 10, 64)
			if err != nil {
				panic(err)
			}
			if !val.IsLength(0, int(i)) {
				return nil, ErrMax
			}
		case "strict-min":
			if opt == "" {
				panic("min requires a number")
			}
			i, err := strconv.ParseInt(fmt.Sprintf("%v", opt), 10, 64)
			if err != nil {
				panic(err)
			}
			if !val.MustLength(int(i), 0) {
				return nil, ErrMin
			}
		case "strict-max":
			if opt == "" {
				panic("max requires a number")
			}
			i, err := strconv.ParseInt(fmt.Sprintf("%v", opt), 10, 64)
			if err != nil {
				panic(err)
			}
			if !val.MustLength(0, int(i)) {
				return nil, ErrMax
			}
		}
	}
	switch format {
	case "string":
		return val.Val(), nil
	case "integer":
		return val.Integer().Val(), nil
	case "float":
		return val.Float().Val(), nil
	case "bool":
		return val.Bool().Val(), nil
	}
	return val, nil
}
