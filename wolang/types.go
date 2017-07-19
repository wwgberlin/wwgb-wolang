package wolang

import (
	"fmt"
)

type(
	DataType interface {
		GetValue() interface{}
	}

	Addable interface {
		DataType
		Add(d Addable) (Addable, error)
	}

	Integer struct {
		value int
	}

	String struct {
		value string
	}

	Array struct {
		value []DataType
	}

	Boolean struct {
		value bool
	}
)

func NewBoolean(b bool) Boolean{
	return Boolean{b}
}

func NewInteger(i int) Integer{
	return Integer{i}
}

func NewString(s string) String{
	return String{s}
}

func (str String) GetValue() interface{} {
	return str.value;
}

func (i Integer) GetValue() interface{} {
	return i.value;
}

func (b Boolean) GetValue() interface{} {
	return b.value;
}

func (arr Array) GetValue() interface{} {
	return arr.value;
}

func (str String) Add(d Addable) (Addable, error) {
	return String{fmt.Sprintf("%v%v", str.value, d.GetValue())}, nil
}

func (i Integer) Add(d Addable) (Addable, error) {
	if di, ok := d.(Integer); !ok {
		return nil, fmt.Errorf("cannot add %v to an integer", d.GetValue())
	} else {
		return Integer{i.value + di.value}, nil
	}
}

