package wolang

import (
	"testing"
)

func TestAddTypeInteger(t *testing.T) {
	zero := Integer{0}
	one := Integer{1}
	if res, err := zero.Add(one); err != nil {
		t.Error("Unexpected error", err);
	} else if res.GetValue().(int) != 1 {
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
