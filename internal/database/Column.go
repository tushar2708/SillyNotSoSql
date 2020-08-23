package database

import (
	"fmt"

	"InMemorySqlDB/internal/types"
	"InMemorySqlDB/internal/verification"
)

type Column struct {
	name  string
	dType DataType
}

type RowValue struct {
	Value interface{}
}

func NewColumn(name string, dType types.DataTypePicker, nullable bool, min, max int) (*Column, error) {

	// Ideally, it should be a Column factory
	if dType == types.INT {
		i, iErr := verification.NewInteger(name, nullable, max, min)
		if iErr != nil {
			return nil, iErr
		}
		return &Column{
			name:  name,
			dType: i,
		}, nil
	}

	if dType == types.STRING {
		i, iErr := verification.NewString(name, nullable, max, min)
		if iErr != nil {
			return nil, iErr
		}
		return &Column{
			name:  name,
			dType: i,
		}, nil
	}

	return nil, fmt.Errorf("invalid datatype")
}
