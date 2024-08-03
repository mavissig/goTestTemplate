package template

import (
	"errors"
)

type ExampleType struct {
	FieldString   string
	FieldInt      int64
	FieldSliceInt []int64
}

func New() *ExampleType {
	return &ExampleType{}
}

func (ps *ExampleType) SetFieldString(str string) error {
	if len(str) < 8 {
		return errors.New("invalid string")
	}
	ps.FieldString = str
	return nil
}
