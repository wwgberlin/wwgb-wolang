package wolang

import (
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		input  string
		result interface{}
	}{{
		`(concat a b cde)`,
		"abcde",
	}, {
		`(+ 1 2 7 7)`,
		17,
		//}, {
		//	`"hi there!"`,
		//	"hi there!",
	}, {
		`27`,
		27,
	}, {
		`(+ 3 4 (+ 5 6)))`,
		18,
	}}

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

func TestEvalErrors(t *testing.T) {
	tests := []struct {
		input        string
		errorMessage string
	}{{
		`(concat a b 17)`,
		"error: '17' is not a string I can concat!",
	}, {
		`(+ 1 2 true)`,
		"error: 'true' is not a number I can add!",
	}, {
		`(minus 3 2)`,
		"error: Function 'minus' is not defined",
	}}

	for _, tt := range tests {

		_, parsed := MustParse(tt.input)
		result, err := Eval(parsed)
		if err == nil {
			t.Errorf("Expected %v to evaluate to an error, got %v instead", tt.input, result)
		} else {
			checkError(t, err, tt.errorMessage)
		}
	}
}

// -----------------------------------------------------------------------------
// misc. helpers:
// -----------------------------------------------------------------------------

func checkError(t *testing.T, err error, prefix string) {
	found := err.Error()
	if !strings.HasPrefix(found, prefix) {
		t.Errorf(
			"Expected:\n\n  | %s\n\n  to match:\n\n  | %s\n",
			strings.Replace(found, "\n", "\n  | ", -1),
			strings.Replace(prefix, "\n", "\n  | ", -1),
		)
	}
}
