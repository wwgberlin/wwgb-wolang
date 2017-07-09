package wolang

import (
	"fmt"
	"strconv"
	"strings"
)

func MustParse(input string) (unparsed string, expr interface{}) {
	u, expr, err := Parse(input)
	if err != nil {
		panic(err.Error())
	}
	return u, expr
}

func Parse(input string) (unparsed string, expr interface{}, err error) {
	// remove all whitespace or newline at the beginning
	input = strings.TrimLeft(input, "\t\r\n ")

	if len(input) == 0 {
		// if nothing's left return empty unparsed
		return unparsed, nil, nil
	} else if input[0] == '(' {
		// if it opens a bracket parse the expression inside
		return parseProcCall(input)
	} else {
		return parseAtom(input)
	}
}

func parseAtom(input string) (unparsed string, expr interface{}, err error) {
	var atom = ""
	for p := 0; p < len(input); p++ {
		if isWhitespace(input[p]) || isEndOfAtom(input[p]) {
			atom = input[:p]
			unparsed = input[p:]
			break
		}
	}

	if atom == "" {
		atom = input
	}

	// bool expressions
	if atom == "TRUE" || atom == "true" {
		return unparsed, true, nil
	} else if atom == "FALSE" || atom == "false" {
		return unparsed, false, nil
	} else {
		// numeric expressions
		val, intErr := strconv.ParseInt(atom, 10, 64)
		if conversionError, ok := intErr.(*strconv.NumError); ok {
			if conversionError.Err == strconv.ErrSyntax {
				valFl, flErr := strconv.ParseFloat(atom, 64)
				if conversionErrorFl, okFl := flErr.(*strconv.NumError); okFl {
					if conversionErrorFl.Err == strconv.ErrSyntax {
						// ...fall out of numeric expr. conditions and fallthrough

					} else {
						return unparsed, float64(0), fmt.Errorf("number out of range: ", atom)
					}
				} else {
					return unparsed, float64(valFl), nil
				}

				// ...fall out of numeric expr. conditions and fallthrough
			} else {
				return unparsed, int(0), fmt.Errorf("number out of range: ", atom)
			}
		} else {
			return unparsed, int(val), nil
		}
	}

	// ...everything else is a string
	return unparsed, string(atom), nil
}

func parseProcCall(input string) (unparsed string, expr []interface{}, err error) {
	expr = []interface{}{}

	// skip opening '('
	input = input[1:]

	// parse input expr
	for p := 0; p < len(input); p++ {

		if isWhitespace(input[p]) {
			continue
		} else if input[p] == ')' {
			return input[p+1:], expr, err
		} else {
			remaining, parsedExpr, err := Parse(input[p:])
			if err != nil {
				return remaining, expr, err
			}

			expr = append(expr, parsedExpr)
			input = remaining
			p = -1 // <-- hacky hack
		}
	}

	return unparsed, expr, fmt.Errorf("syntax error, missing expected ')' before end of input")
}

func isEndOfAtom(c byte) bool {
	return c == '\n' || c == ')'
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r'
}
