package vision

import (
	"fmt"
	"strconv"
)

// ParseInteger parse value and return Integer object
func ParseInteger(arg interface{}) (Integer, Error) {
	switch obj := arg.(type) {
	case string:
		i, err := strconv.ParseInt(obj, 10, 64)
		if err != nil {
			return 0, ErrInvalid
		}
		return Integer(i), nil
	case bool:
		if obj {
			return Integer(1), nil
		}
		return Integer(0), nil
	case int:
		return Integer(obj), nil
	case int8:
		return Integer(obj), nil
	case int32:
		return Integer(obj), nil
	case int64:
		return Integer(obj), nil
	case uint:
		return Integer(obj), nil
	case uint8:
		return Integer(obj), nil
	case uint32:
		return Integer(obj), nil
	case uint64:
		return Integer(obj), nil
	case float32:
		return Integer(obj), nil
	case float64:
		return Integer(obj), nil
	case complex64, complex128:
		i, err := strconv.ParseInt(fmt.Sprintf("%f", obj), 10, 64)
		if err != nil {
			return 0, ErrInvalid
		}
		return Integer(i), nil
	}
	return 0, ErrUnsupported
}

// IsPositive checks if the integer number is positive
func (i Integer) IsPositive() bool {
	return i >= 0
}

// IsNegative checks if the integer number is negative
func (i Integer) IsNegative() bool {
	return i < 0
}

// IsBetween checks if value is between (a) and (b)
func (i Integer) IsBetween(min, max int) bool {
	return int(i) >= min && int(i) <= max
}

// IsLatitude checks if value is a valid latitude
func (i Integer) IsLatitude() bool {
	s := fmt.Sprintf("%d", i)
	return regexLatitude.MatchString(s)
}

// IsLongitude checks if value is a valid longitude
func (i Integer) IsLongitude() bool {
	s := fmt.Sprintf("%d", i)
	return regexLongitude.MatchString(s)
}

// IsUSPhoneNumber checks if the value is a valid us phone number
func (i Integer) IsUSPhoneNumber() bool {
	s := fmt.Sprintf("%d", i)
	return regexUSPhoneNumber.MatchString(s)
}

// String returns the string type value of the integer
func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

// Bool returns the bool type value of the integer
func (i Integer) Bool() Bool {
	if i == 1 {
		return Bool(true)
	}
	return Bool(false)
}

// Float returns the float type value of the integer
func (i Integer) Float() Float {
	return Float(i)
}

// Val returns the build-in type value
func (i Integer) Val() int {
	return int(i)
}
