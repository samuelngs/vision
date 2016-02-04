package vision

import (
	"fmt"
	"strconv"
)

// ParseFloat parse value and return Float object
func ParseFloat(arg interface{}) (Float, Error) {
	switch obj := arg.(type) {
	case string:
		i, err := strconv.ParseFloat(obj, 64)
		if err != nil {
			return 0, ErrInvalid
		}
		return Float(i), nil
	case bool:
		if obj {
			return Float(1.0), nil
		}
		return Float(0.0), nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		i, err := strconv.ParseFloat(fmt.Sprintf("%d", obj), 64)
		if err != nil {
			return 0, ErrInvalid
		}
		return Float(i), nil
	case float32, float64, complex64:
		i, err := strconv.ParseFloat(fmt.Sprintf("%f", obj), 64)
		if err != nil {
			return 0, ErrInvalid
		}
		return Float(i), nil
	}
	return 0, nil
}

// IsPositive checks if the integer number is positive
func (f Float) IsPositive() bool {
	return f >= 0.0
}

// IsNegative checks if the integer number is negative
func (f Float) IsNegative() bool {
	return f < 0.0
}

// IsBetween checks if value is between (a) and (b)
func (f Float) IsBetween(min, max float64) bool {
	return float64(f) >= min && float64(f) <= max
}

// IsLatitude checks if value is a valid latitude
func (f Float) IsLatitude() bool {
	s := fmt.Sprintf("%f", f)
	return regexLatitude.MatchString(s)
}

// IsLongitude checks if value is a valid longitude
func (f Float) IsLongitude() bool {
	s := fmt.Sprintf("%f", f)
	return regexLongitude.MatchString(s)
}

// String returns the string type value of the float number
func (f Float) String() string {
	return fmt.Sprintf("%f", f)
}

// Bool returns the bool type value of the float number
func (f Float) Bool() Bool {
	if f == 1.0 {
		return Bool(true)
	}
	return Bool(false)
}

// Integer returns the float type value of the float number
func (f Float) Integer() Integer {
	return Integer(f)
}

// Val returns the build-in type value
func (f Float) Val() float64 {
	return float64(f)
}
