package payroll24

import (
	"strconv"
	"time"
)

// StringInt represents an integer value that comes as a string from the API
type StringInt string

// UnmarshalJSON implements json.Unmarshaler interface
func (s *StringInt) UnmarshalJSON(data []byte) error {
	// Remove quotes if present
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	*s = StringInt(str)
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (s StringInt) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(s) + `"`), nil
}

// Int returns the integer value
func (s StringInt) Int() int {
	i, _ := strconv.Atoi(string(s))
	return i
}

// IntPtr returns a pointer to the integer value, nil if empty
func (s StringInt) IntPtr() *int {
	if s == "" {
		return nil
	}
	i := s.Int()
	return &i
}

// String implements the Stringer interface
func (s StringInt) String() string {
	return string(s)
}

// StringFloat represents a float value that comes as a string from the API
type StringFloat string

// UnmarshalJSON implements json.Unmarshaler interface
func (s *StringFloat) UnmarshalJSON(data []byte) error {
	// Remove quotes if present
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	*s = StringFloat(str)
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (s StringFloat) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(s) + `"`), nil
}

// Float64 returns the float64 value
func (s StringFloat) Float64() float64 {
	f, _ := strconv.ParseFloat(string(s), 64)
	return f
}

// Float64Ptr returns a pointer to the float64 value, nil if empty
func (s StringFloat) Float64Ptr() *float64 {
	if s == "" {
		return nil
	}
	f := s.Float64()
	return &f
}

// String implements the Stringer interface
func (s StringFloat) String() string {
	return string(s)
}

// StringBool represents a boolean value that comes as "0"/"1" string from the API
type StringBool string

// UnmarshalJSON implements json.Unmarshaler interface
func (s *StringBool) UnmarshalJSON(data []byte) error {
	// Remove quotes if present
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	*s = StringBool(str)
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (s StringBool) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(s) + `"`), nil
}

// Bool returns the boolean value (true for "1", false for "0" or anything else)
func (s StringBool) Bool() bool {
	return string(s) == "1"
}

// BoolPtr returns a pointer to the boolean value, nil if empty
func (s StringBool) BoolPtr() *bool {
	if s == "" {
		return nil
	}
	b := s.Bool()
	return &b
}

// String implements the Stringer interface
func (s StringBool) String() string {
	return string(s)
}

// StringDate represents a date value that comes as a string from the API
type StringDate string

// UnmarshalJSON implements json.Unmarshaler interface
func (s *StringDate) UnmarshalJSON(data []byte) error {
	// Remove quotes if present
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	*s = StringDate(str)
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (s StringDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(s) + `"`), nil
}

// Time returns the time.Time value, assumes YYYY-MM-DD format
func (s StringDate) Time() time.Time {
	t, _ := time.Parse("2006-01-02", string(s))
	return t
}

// TimePtr returns a pointer to the time.Time value, nil if empty
func (s StringDate) TimePtr() *time.Time {
	if s == "" {
		return nil
	}
	t := s.Time()
	return &t
}

// String implements the Stringer interface
func (s StringDate) String() string {
	return string(s)
}
