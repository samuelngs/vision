package vision

var (
	// ErrEmpty is the error returned when the value is empty
	ErrEmpty = ValueError{"value is empty"}
	// ErrZeroValue is the error returned when variable has zero valud
	// and nonzero was specified
	ErrZeroValue = ValueError{"zero value"}
	// ErrMin is the error returned when variable is less than mininum
	// value specified
	ErrMin = ValueError{"less than min"}
	// ErrMax is the error returned when variable is more than
	// maximum specified
	ErrMax = ValueError{"greater than max"}
	// ErrLen is the error returned when length is not equal to
	// param specified
	ErrLen = ValueError{"invalid length"}
	// ErrRegexp is the error returned when the value does not
	// match the provided regular expression parameter
	ErrRegexp = ValueError{"regular expression mismatch"}
	// ErrUnsupported is the error error returned when a validation rule
	// is used with an unsupported variable type
	ErrUnsupported = ValueError{"unsupported type"}
	// ErrBadParameter is the error returned when an invalid parameter
	ErrBadParameter = ValueError{"bad parameter"}
	// ErrUnknownTag is the error returned when an unknown tag is found
	ErrUnknownTag = ValueError{"unknown tag"}
	// ErrInvalid is the error returned when variable is invalid
	ErrInvalid = ValueError{"invalid value"}
	// ErrType is the error returned when data is invalid type
	ErrType = ValueError{"invalid email"}
)

type (
	// Error Interface
	Error interface {
		Error() string
		MarshalText() ([]byte, error)
	}

	// ValueError is an vision value error object
	ValueError struct {
		ErrorString string
	}
)

// Error implements the error interface.
func (t ValueError) Error() string {
	return t.ErrorString
}

// MarshalText implements the TextMarshaller
func (t ValueError) MarshalText() ([]byte, error) {
	return []byte(t.ErrorString), nil
}
