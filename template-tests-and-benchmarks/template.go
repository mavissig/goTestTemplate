package template_tests_and_benchmarks

import "errors"

type ExampleType struct {
	FieldString   string
	FieldInt      int64
	FieldSliceInt []int64
}

func New() *ExampleType {
	return &ExampleType{
		FieldSliceInt: make([]int64, 0),
	}
}

func (ps *ExampleType) AddFieldIntToFieldSliceInt() error {
	if ps.FieldInt < 200 || ps.FieldInt > 299 {
		return errors.New("FieldInt is not equal to 200")
	}
	ps.FieldSliceInt = append(ps.FieldSliceInt, ps.FieldInt)
	return nil
}

func (ps *ExampleType) SetFieldSliceInt(nums []int64) error {
	if len(ps.FieldSliceInt) != 0 {
		return errors.New("ExampleType.FieldSliceInt is not empty")
	} else if len(nums) < 2 {
		return errors.New("incorrect size of transmitted slice")
	}
	ps.FieldSliceInt = nums
	return nil
}

func (ps *ExampleType) GetFieldSliceInt() []int64 {
	return ps.FieldSliceInt
}

func (ps *ExampleType) SetFieldString(str string) error {
	if len(str) < 8 {
		return errors.New("invalid string")
	}
	ps.FieldString = str
	return nil
}
