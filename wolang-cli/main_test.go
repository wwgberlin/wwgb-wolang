package main

import (
	"testing"

	. "github.com/wwgberlin/wwgb-wolang/wolang"
)

func TestInvalidLargerThanOneAttribute(t *testing.T) {
	if _, err := largerThan([]DataType{NewInteger(1)}); err == nil {
		t.Error("Expected to receive an error when receiving only one argument with nothing to compare it to");
	} else if err.Error() != ERR_ARGS_TOO_FEW {
		t.Errorf("Expected to return error %s", ERR_ARGS_TOO_FEW);
	}
}

func TestInvalidLargerThanTooManyAttributes(t *testing.T) {
	if _, err := largerThan([]DataType{NewInteger(1), NewInteger(2), NewInteger(3)}); err == nil {
		t.Errorf("Expected to receive error %s", ERR_ARGS_TOO_MANY);
	} else if err.Error() != ERR_ARGS_TOO_MANY {
		t.Errorf("Expected to return error %s", ERR_ARGS_TOO_MANY);
	}
}

func TestInvalidLargerThanNilAttributes(t *testing.T) {
	if _, err := largerThan([]DataType{nil, nil}); err == nil {
		t.Error("Expected to return error %s", ERR_ARGS_INVALID);
	} else if err.Error() != ERR_ARGS_INVALID {
		t.Errorf("Expected to return error %s", ERR_ARGS_INVALID);
	}
}

func TestInvalidLargerThanStringAttributes(t *testing.T) {
	if _, err := largerThan([]DataType{NewString("1"), NewString("2")}); err == nil {
		t.Error("Expected to return error %s", ERR_ARGS_INVALID);
	} else if err.Error() != ERR_ARGS_INVALID {
		t.Errorf("Expected to return error %s", ERR_ARGS_INVALID);
	}
}

func TestInvalidEqualsOneAttribute(t *testing.T) {
	if _, err := equals([]DataType{NewInteger(1)}); err == nil {
		t.Error("Expected to receive an error when receiving only one argument with nothing to compare it to");
	} else if err.Error() != ERR_ARGS_TOO_FEW {
		t.Errorf("Expected to return error %s", ERR_ARGS_TOO_FEW);
	}
}

func TestInvalidEqualsTooManyAttributes(t *testing.T) {
	if _, err := equals([]DataType{NewInteger(1), NewInteger(2), NewInteger(3)}); err == nil {
		t.Errorf("Expected to receive error %s", ERR_ARGS_TOO_MANY);
	} else if err.Error() != ERR_ARGS_TOO_MANY {
		t.Errorf("Expected to return error %s", ERR_ARGS_TOO_MANY);
	}
}

func TestInvalidEqualsNilAttributes(t *testing.T) {
	if _, err := equals([]DataType{nil, nil}); err == nil {
		t.Error("Expected to return error %s", ERR_ARGS_INVALID);
	} else if err.Error() != ERR_ARGS_INVALID {
		t.Errorf("Expected to return error %s", ERR_ARGS_INVALID);
	}
}

func TestInvalidEqualsStringAttributes(t *testing.T) {
	if _, err := equals([]DataType{NewString("2"), NewString("1")}); err == nil {
		t.Error("Expected to return error %s", ERR_ARGS_INVALID);
	} else if err.Error() != ERR_ARGS_INVALID {
		t.Errorf("Expected to return error %s", ERR_ARGS_INVALID);
	}
}

func TestInvalidSmallerThanOneAttribute(t *testing.T) {
	if _, err := smallerThan([]DataType{NewInteger(1)}); err == nil {
		t.Error("Expected to receive an error when receiving only one argument with nothing to compare it to");
	} else if err.Error() != ERR_ARGS_TOO_FEW {
		t.Errorf("Expected to return error %s", ERR_ARGS_TOO_FEW);
	}
}

func TestInvalidSmallerThanTooManyAttributes(t *testing.T) {
	if _, err := smallerThan([]DataType{NewInteger(1), NewInteger(2), NewInteger(3)}); err == nil {
		t.Errorf("Expected to receive error %s", ERR_ARGS_TOO_MANY);
	} else if err.Error() != ERR_ARGS_TOO_MANY {
		t.Errorf("Expected to return error %s", ERR_ARGS_TOO_MANY);
	}
}

func TestInvalidSmallerThanNilAttributes(t *testing.T) {
	if _, err := smallerThan([]DataType{nil, nil}); err == nil {
		t.Error("Expected to return error %s", ERR_ARGS_INVALID);
	} else if err.Error() != ERR_ARGS_INVALID {
		t.Errorf("Expected to return error %s", ERR_ARGS_INVALID);
	}
}

func TestInvalidSmallerThanStringAttributes(t *testing.T) {
	if _, err := smallerThan([]DataType{NewString("2"), NewString("1")}); err == nil {
		t.Error("Expected to return error %s", ERR_ARGS_INVALID);
	} else if err.Error() != ERR_ARGS_INVALID {
		t.Errorf("Expected to return error %s", ERR_ARGS_INVALID);
	}
}

type(
	testCase struct {
		left        int
		right       int
		expectedRes bool
	}
)

func TestOperators(t *testing.T) {
	funcs := map[string]func(terms []DataType) (result DataType, err error){">": largerThan, "<": smallerThan, "==": equals}
	toTest := map[string][]testCase{
		">": []testCase{
			{2, 1, true},
			{1, 2, false},
			{1, 1, false},
		},
		"<": []testCase{
			{1, 2, true},
			{1, 1, false},
			{2, 1, false},
		},
		"==": []testCase{
			{1, 1, true},
			{2, 1, false},
			{1, 2, false},
		},
	}

	for operator, testCases := range toTest {
		for _, testCase := range testCases {
			f := funcs[operator]
			res, err := f([]DataType{NewInteger(testCase.left), NewInteger(testCase.right)})
			if err != nil {
				t.Errorf("Unexpected error returned when evaluating %v %s %v", testCase.left, operator, testCase.right);
			} else if b, ok := res.GetValue().(bool); !ok {
				t.Errorf("Unexpected type returned when evaluating %v %s %v", testCase.left, operator, testCase.right);
			} else if testCase.expectedRes != b {
				t.Error("%v %s %v is expected to be: %v", testCase.left, operator, testCase.right, testCase.expectedRes);
			}
		}
	}
}
