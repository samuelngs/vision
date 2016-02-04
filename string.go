package vision

import (
	"fmt"
	"strings"
)

// ParseString parse string value and return String object
func ParseString(arg interface{}) (String, Error) {
	switch obj := arg.(type) {
	case string:
		return String(obj), nil
	case bool:
		return String(fmt.Sprintf("%t", obj)), nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		return String(fmt.Sprintf("%d", obj)), nil
	case float32, float64, complex64:
		return String(fmt.Sprintf("%f", obj)), nil
	}
	return "", ErrUnsupported
}

// String returns the string value of String instance
func (str String) String() string {
	return fmt.Sprintf("%v", str)
}

// Integer returns the integer value of the string (has to be an integer number)
func (str String) Integer() Integer {
	if !str.IsInteger() {
		return 0
	}
	if i, err := ParseInteger(str.String()); err == nil {
		return i
	}
	return 0
}

// Float returns the float value of the string (has to be an float number)
func (str String) Float() Float {
	if !str.IsFloat() {
		return 0.0
	}
	if f, err := ParseFloat(str.String()); err == nil {
		return f
	}
	return 0.0
}

// Bool returns the bool value of the string (has to be an bool value, otherwise always false)
func (str String) Bool() Bool {
	s := strings.Trim(strings.ToLower(str.String()), Trim)
	if s == "true" {
		return Bool(true)
	}
	return Bool(false)
}

// IsEmpty checks if string is empty
func (str String) IsEmpty() bool {
	return len(str) == 0
}

// MustEmpty checks if string is absolute empty (trim all spaces and non-alphabetic characters
func (str String) MustEmpty() bool {
	return len(strings.Trim(str.String(), Trim)) == 0
}

// IsLength checks if string length is valid
func (str String) IsLength(min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// MustLength checks if trimed string is match its length
func (str String) MustLength(min, max int) bool {
	count := len(strings.Trim(str.String(), Trim))
	return count >= min && count <= max
}

// IsEmail checks if string is a valid email address
func (str String) IsEmail() bool {
	return regexEmail.MatchString(str.String())
}

// IsURL checks if string is a valid URL
func (str String) IsURL() bool {
	return regexURL.MatchString(str.String())
}

// IsLatitude checks if value is a valid latitude
func (str String) IsLatitude() bool {
	return regexLatitude.MatchString(str.String())
}

// IsLongitude checks if value is a valid longitude
func (str String) IsLongitude() bool {
	return regexLongitude.MatchString(str.String())
}

// IsUUID checks if string is in UUID format
func (str String) IsUUID() bool {
	return regexUUID.MatchString(str.String())
}

// IsUUID3 checks if string is in UUID3 format
func (str String) IsUUID3() bool {
	return regexUUID3.MatchString(str.String())
}

// IsUUID4 checks if string is in UUID4 format
func (str String) IsUUID4() bool {
	return regexUUID4.MatchString(str.String())
}

// IsUUID5 checks if string is in UUID5 format
func (str String) IsUUID5() bool {
	return regexUUID5.MatchString(str.String())
}

// IsBase64 checks if string is in Base64 format
func (str String) IsBase64() bool {
	return regexBase64.MatchString(str.String())
}

// IsDataURI checks if string is in DataURI format
func (str String) IsDataURI() bool {
	return regexDataURI.MatchString(str.String())
}

// IsHexColor checks if string is an hex color value
func (str String) IsHexColor() bool {
	return regexHexColor.MatchString(str.String())
}

// IsRGBColor checks if string is a valid RGB color value
func (str String) IsRGBColor() bool {
	return regexRGBColor.MatchString(str.String())
}

// IsCreditCard checks if string a credit card number vale
func (str String) IsCreditCard() bool {
	return regexCreditCard.MatchString(str.String())
}

// IsNumberic checks if string is actually a number
func (str String) IsNumberic() bool {
	return regexNumeric.MatchString(str.String())
}

// IsInteger checks if string is a integer
func (str String) IsInteger() bool {
	return regexInt.MatchString(str.String())
}

// IsFloat checks if string is a float number
func (str String) IsFloat() bool {
	return regexFloat.MatchString(str.String())
}

// Val returns the build-in type value
func (str String) Val() string {
	return str.String()
}
