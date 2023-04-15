package util

import (
	"strconv"
	"strings"
)

type IVariable interface {
	DefaultString(defaultValue string) string
	DefaultBool(defaultValue bool) bool
	DefaultInt(defaultValue int) int
	DefaultStrings(defaultValue []string) []string
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

func (v *variable) DefaultInt(defaultValue int) int {
	if len(v.Value) == 0 {
		return defaultValue
	}

	value, err := strconv.Atoi(v.Value)
	if err != nil {
		return defaultValue
	}

	return value
}

func (v *variable) DefaultStrings(defaultValue []string) []string {
	if len(v.Value) == 0 {
		return []string{}
	}
	return strings.Split(v.Value, ",")
}

func NewVariable(value string) IVariable {
	return &variable{
		Value: value,
	}
}
