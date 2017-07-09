package wolang

import (
	"testing"
)

func TestExtension(t *testing.T) {
	tests := []struct {
		input  string
		result interface{}
	}{{
		`(always99)`,
		99,
	}}

	var extensions []ExtFuncDef = []ExtFuncDef{{
		"always99",
		func(terms []interface{}) (result interface{}, err error) {
			return 99, nil
		}},
	}

	for _, ex := range extensions {
		RegExtFunc(ex)
	}

	for i, tt := range tests {

		_, parsed := MustParse(tt.input)
		result, err := Eval(parsed)
		if err != nil {
			t.Errorf("Error evaluating #%d: %v %v", i, tt.input, err)
		} else {
			if result != tt.result {
				t.Error("Expected", result, "to equal", tt.result, "when evaluating", tt.input)
			}
		}
	}
}
