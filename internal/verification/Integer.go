package verification

import (
	"fmt"
)

type Integer struct {
	name      string
	validator DataValidator
	//rules SchemaRules
}

func NewInteger(name string, nullable bool, max, min int) (*Integer, error) {

	if max > INT_MAX || min < INT_MIN {
		return nil, fmt.Errorf("can't support integers beyond the length of [%d,%d]", INT_MIN, INT_MAX)
	}

	return &Integer{
		name: name,
		validator: &IntegerValidator{
			rules: &IntegerSchema{
				maxValue: max,
				minValue: min,
				nullable: nullable,
			},
		},
	}, nil
}

type IntegerSchema struct {
	maxValue int
	minValue int
	nullable bool
}

func (i IntegerSchema) IsNullable() bool {
	return i.nullable
}

func (i IntegerSchema) MaxValue() int {
	return i.maxValue
}

func (i IntegerSchema) MinValue() int {
	return i.minValue
}

type IntegerValidator struct {
	rules SchemaRules
}

func (i IntegerValidator) IsTypeValid(value interface{}) (bool, error) {
	_, ok := value.(int)
	if !ok {
		return false, fmt.Errorf("expected integer, found invalid type for [%v]", value)
	}

	return true, nil
}

func (i IntegerValidator) IsNullabilityValid(value interface{}) (bool, error) {

	//true == true
	// false == false
	if (value == nil) && !i.rules.IsNullable() {
		return false, fmt.Errorf("value must not be null [%v]", value)
	}

	return true, nil
}

func (i IntegerValidator) IsMaxValid(value interface{}) (bool, error) {

	v := value.(int)

	if v < i.rules.MinValue() {
		return false, fmt.Errorf("value can't be less than %d", i.rules.MinValue())
	}

	return true, nil
}

func (i IntegerValidator) IsMinValid(value interface{}) (bool, error) {

	v := value.(int)

	if v > i.rules.MaxValue() {
		return false, fmt.Errorf("value can't be greater than %d", i.rules.MaxValue())
	}

	return true, nil
}

func (i Integer) IsValidValue(value interface{}) (bool, error) {

	if ok, nErr := i.validator.IsNullabilityValid(value); !ok {
		return ok, fmt.Errorf("for value %s, %v", i.name, nErr)
	}

	if vOk, vErr := i.validator.IsTypeValid(value); !vOk {
		return vOk, fmt.Errorf("for value %s, %v", i.name, vErr)
	}

	if v, err := i.validator.IsMinValid(value); !v {
		return v, fmt.Errorf("for value %s, %v", i.name, err)
	}

	if v, err := i.validator.IsMaxValid(value); !v {
		return v, fmt.Errorf("for value %s, %v", i.name, err)
	}

	return true, nil
}
