package wolang

import (
	"testing"
)

func TestAddTypeInteger(t *testing.T) {
	zero := Integer{0}
	one := Integer{1}
	if res, err := zero.Add(one); err != nil {
		t.Error("Unexpected error", err);
	} else if res.GetValue().(int64) != 1 {
		t.Error("Expected res to be", 1, res.GetValue());
	}
}

func TestAddTypeString(t *testing.T) {
	zero := String{"0"}
	one := String{"1"}
	if res, err := zero.Add(one); err != nil {
		t.Error("Unexpected error", err);
	} else if res.GetValue().(string) != "01" {
		t.Error("Expected res to be", "01", res.GetValue());
	}
}

func TestAddTypeFloat(t *testing.T) {
	zero := Float{0}
	one := Float{1}
	if res, err := zero.Add(one); err != nil {
		t.Error("Unexpected error", err);
	} else if res.GetValue().(float64) != 1 {
		t.Error("Expected res to be", 1, res.GetValue());
	}
}

func TestArrayPush(t *testing.T) {
	zero := Integer{0}
	one := Integer{1}
	arrzero := Array{[]DataType{zero}}
	if res, err := arrzero.Push(one); err != nil {
		t.Error("Unexpected error", err);
	} else {
		if res.GetValue().([]DataType)[0].GetValue().(int64) != 0 {
			t.Error("Expected res[0] to be", 0, res.GetValue());
		}
		if res.GetValue().([]DataType)[1].GetValue().(int64) != 1 {
			t.Error("Expected res[1] to be", 1, res.GetValue());
		}
	}
}
