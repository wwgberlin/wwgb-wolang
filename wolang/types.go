package wolang

import (
	"fmt"
)

type(
	Addable interface {
		DataType
		Add(d Addable) (Addable, error)
	}

	DataType interface {
		GetValue() interface{}
	}

	Integer struct {
		value int64
	}

	Float struct {
		value float64
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

func (str String) GetValue() interface{} {
	return str.value;
}

func (i Integer) GetValue() interface{} {
	return i.value;
}

func (f Float) GetValue() interface{} {
	return f.value;
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

func (f Float) Add(d Addable) (Addable, error) {
	if df, ok := d.(Float); !ok {
		if di, ok := d.(Integer); !ok {
			return nil, fmt.Errorf("cannot add %v to a float", d.GetValue())
		}else{
			return Float{f.value + float64(di.value)}, nil
		}
	} else {
		return Float{f.value + df.value}, nil
	}
}
