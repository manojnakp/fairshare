package config

import "flag"

// String is a custom flag type that represents any string value.
type String string

// String implements flag.Value. String is nil-safe.
func (s *String) String() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

// Set implements flag.Value.
func (s *String) Set(value string) error {
	*s = String(value)
	return nil
}

// Get implements flag.Getter. Returns a value of type `string`,
// or `nil` in case of a nil receiver.
func (s *String) Get() any {
	if s == nil {
		return nil
	}
	return string(*s)
}

// Assert interface satisfaction.
var _ flag.Getter = (*String)(nil)
