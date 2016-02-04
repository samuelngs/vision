package vision

import (
	"fmt"
	"strconv"
)

// ParseInteger parse value and return Integer object
func ParseInteger(arg interface{}) (Integer, Error) {
	switch obj := arg.(type) {
	case string:
		s, err := ParseString(obj)
		if err != nil {
			return 0, ErrInvalid
		}
		if !s.IsInteger() {
			return 0, ErrInvalid
		}
		return s.Integer(), ErrInvalid
	case bool:
		if obj {
			return Integer(1), nil
		}
		return Integer(0), nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		i, err := strconv.ParseInt(fmt.Sprintf("%d", obj), 10, 64)
		if err != nil {
			return 0, ErrInvalid
		}
		return Integer(i), nil
	case float32, float64, complex64:
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
