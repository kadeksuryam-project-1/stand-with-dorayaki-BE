package util

import "strconv"

type IVariable interface {
	DefaultString(defaultValue string) string
	DefaultBool(defaultValue bool) bool
}

type variable struct {
	Value string
}

func (v *variable) DefaultString(defaultValue string) string {
	if len(v.Value) == 0 {
		return defaultValue
	}
	return v.Value
}

func (v *variable) DefaultBool(defaultValue bool) bool {
	if len(v.Value) == 0 {
		return defaultValue
	}

	value, err := strconv.ParseBool(v.Value)
	if err != nil {
		return defaultValue
	}

	return value
}

func NewVariable(value string) IVariable {
	return &variable{
		Value: value,
	}
}
