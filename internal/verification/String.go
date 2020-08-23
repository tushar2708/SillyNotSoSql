package verification

import (
	"fmt"
)

type StringSchema struct {
	maxValue int
	minValue int
	nullable bool
}

func (i StringSchema) IsNullable() bool {
	return i.nullable
}

func (i StringSchema) MaxValue() int {
	return i.maxValue
}

func (i StringSchema) MinValue() int {
	return i.minValue
}

type String struct {
	name      string
	validator DataValidator
}

type StringValidator struct {
	rules SchemaRules
}

func NewString(name string, nullable bool, max, min int) (*String, error) {

	if max > STRING_MAX || min < STRING_MIN {
		return nil, fmt.Errorf("can't support string beyond the length of %d", STRING_MAX)
	}

	return &String{
		name: name,
		validator: &StringValidator{
			rules: &StringSchema{
				maxValue: max,
				minValue: min,
				nullable: nullable,
			},
		},
	}, nil
}

func (i StringValidator) IsTypeValid(value interface{}) (bool, error) {
	_, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("expected String, found invalid type for [%v]", value)
	}

	return true, nil
}

func (i StringValidator) IsNullabilityValid(value interface{}) (bool, error) {

	//true == true
	// false == false
	if (value == nil) && !i.rules.IsNullable() {
		return false, fmt.Errorf("value must not be null [%v]", value)
	}

	return true, nil
}

func (i StringValidator) IsMaxValid(value interface{}) (bool, error) {

	v := value.(string)

	if len(v) < i.rules.MinValue() {
		return false, fmt.Errorf("value can't be less than %d", i.rules.MinValue())
	}

	return true, nil
}

func (i StringValidator) IsMinValid(value interface{}) (bool, error) {

	v := value.(string)

	if len(v) > i.rules.MaxValue() {
		return false, fmt.Errorf("value can't be greater than %d", i.rules.MaxValue())
	}

	return true, nil
}

func (i String) IsValidValue(value interface{}) (bool, error) {

	if ok, nErr := i.validator.IsNullabilityValid(value); !ok {
		return ok, fmt.Errorf("for value %s, %v", i.name, nErr)
	}

	if v, err := i.validator.IsTypeValid(value); !v {
		return v, fmt.Errorf("for value %s, %v", i.name, err)
	}

	if v, err := i.validator.IsMinValid(value); !v {
		return v, fmt.Errorf("for value %s, %v", i.name, err)
	}

	if v, err := i.validator.IsMaxValid(value); !v {
		return v, fmt.Errorf("for value %s, %v", i.name, err)
	}

	return true, nil
}
