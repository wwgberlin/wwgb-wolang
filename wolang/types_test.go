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
	first := Float{1.7}
	second := Float{1.3}
	if res, err := first.Add(second); err != nil {
		t.Error("Unexpected error", err);
	} else if res.GetValue().(float64) != 3.0 {
		t.Error("Expected res to be", 3, res.GetValue());
	}
}

