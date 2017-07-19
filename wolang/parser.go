package wolang

import (
	"fmt"
	"regexp"
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

func Parse(input string) (unparsed string, expr DataType, err error) {
	// remove all whitespace or newline at the beginning
	input = strings.TrimLeft(input, "\t\r\n ")

	if len(input) == 0 {
		// if nothing's left return empty unparsed
		return unparsed, nil, nil
	} else if input[0] == '(' {
		// if it opens a bracket parse the expression inside
		unparsed, expr, err := parseProcCall(input)
		return unparsed, Array{expr}, err
	} else if input[0] == '"' {
		return parseDoubleQuotedString(input)
	} else {
		return parseAtom(input)
	}
}

func parseDoubleQuotedString(input string) (unparsed string, expr DataType, err error) {
	// skip opening '"'
	input = input[1:]

	// check for empty double quoted string
	if input == "\"" {
		return "", String{""}, nil
	}

	var dqStr string
	for p := 0; p < len(input); p++ {
		if input[p] == '"' {
			dqStr = input[:p]
			unparsed = input[p + 1:]
			break
		}

		// allow to escape a double quote
		if input[p] == '\\' {
			if p + 1 == len(input) {
				return unparsed, expr, fmt.Errorf(
					"illegal escape sequence at the end of %s", input)
			}

			if input[p + 1] == '"' || input[p + 1] == '\\' {
				// skip the escaping backslash
				input = input[:p] + input[p + 1:]

				if p == len(input) {
					return unparsed, expr, fmt.Errorf(
						"illegal escape sequence at the end of %s", input)
				}
			}
		}
	}

	if len(dqStr) == 0 {
		return "", expr, fmt.Errorf("unterminated double-quoted string: %s", input)
	}

	return unparsed, String{dqStr}, nil
}

func getInteger(atom string) (int, error) {
	if val, err := strconv.ParseInt(atom, 10, 64); err != nil {
		return 0, fmt.Errorf("number out of range: ", atom)
	} else {
		return int(val), nil
	}
}

var regexInteger *regexp.Regexp = regexp.MustCompile(`^[-+]?[0-9]+$`)

func parseAtom(input string) (unparsed string, expr DataType, err error) {
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
		return unparsed, NewBoolean(true), nil
	}
	if atom == "FALSE" || atom == "false" {
		return unparsed, NewBoolean(false), nil
	}

	// integer
	if isInteger := regexInteger.MatchString(atom); isInteger {
		val, err := getInteger(atom)
		return unparsed, NewInteger(val), err
	}

	// ...everything else is a string
	return unparsed, NewString(string(atom)), nil
}

func parseProcCall(input string) (unparsed string, expr []DataType, err error) {
	expr = []DataType{}

	// skip opening '('
	input = input[1:]

	// parse input expr
	for p := 0; p < len(input); p++ {

		if isWhitespace(input[p]) {
			continue
		} else if input[p] == ')' {
			return input[p + 1:], expr, err
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
