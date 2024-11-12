package validation

import (
	"fmt"
	"reflect"
)

type Input struct {
	Value  interface{}
	Errors []error
}

func I(value interface{}) *Input {
	return &Input{Value: value}
}

func (input *Input) Required() *Input {

	return input
}

func (input *Input) Min(value int) *Input {
	if reflect.TypeOf(input.Value).Name() == "string" && len(input.Value.(string)) < value {
		input.Errors = append(input.Errors, fmt.Errorf("value must be at least %d characters long", value))
	}
	return input
}

func (input *Input) Max(value int) *Input {
	if reflect.TypeOf(input.Value).Name() == "string" && len(input.Value.(string)) > value {
		input.Errors = append(input.Errors, fmt.Errorf("value must be at least %d characters long", value))
	}
	return input
}

func (input *Input) Validate() *Input {

	return input
}
