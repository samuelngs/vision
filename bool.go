package vision

import "strings"

// ParseBool parse value and return Integer object
func ParseBool(arg interface{}) (Bool, Error) {
	switch obj := arg.(type) {
	case string:
		str := strings.Trim(strings.ToLower(obj), Trim)
		if str == "true" {
			return Bool(true), nil
		} else if str == "false" {
			return Bool(false), nil
		}
		return Bool(false), ErrInvalid
	case bool:
		return Bool(obj), nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		switch obj {
		case 0:
			return Bool(false), nil
		case 1:
			return Bool(true), nil
		}
		return Bool(false), ErrInvalid
	case float32, float64, complex64:
		switch obj {
		case 0.0:
			return Bool(false), nil
		case 1.0:
			return Bool(true), nil
		}
		return Bool(false), ErrInvalid
	}
	return false, ErrInvalid
}

// IsTrue checks if the bool value is "true"
func (b Bool) IsTrue() bool {
	return b == true
}

// IsFalse checks if the bool value is "false"
func (b Bool) IsFalse() bool {
	return b == false
}

// Bool returns string type value
func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

// Integer returns the interger value
func (b Bool) Integer() Integer {
	if b {
		return Integer(1)
	}
	return Integer(0)
}

// Float returns the float value
func (b Bool) Float() Float {
	if b {
		return Float(1.0)
	}
	return Float(0.0)
}

// Val returns the build-in type value
func (b Bool) Val() bool {
	return bool(b)
}
