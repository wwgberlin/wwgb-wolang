package wolang

import (
	"fmt"
)

func Eval(expr interface{}) (result interface{}, err error) {
	switch v := expr.(type) {
	case string:
		return v, nil
	case int:
		return v, nil
	case []interface{}:
		return evalFCall(v)
	}
	return nil, fmt.Errorf("error: Illegal expression type %T", expr)
}

func evalFCall(expr []interface{}) (interface{}, error) {
	funcName, isFName := expr[0].(string)
	if !isFName {
		return nil, fmt.Errorf("error: Illegal call. %v is not a function name", expr[0])
	}
	arguments := expr[1:]

	for ind, arg := range arguments {
		if argarray, ok := arg.([]interface{}); ok {
			nestresult, err := evalFCall(argarray)
			if err != nil {
				return nil, fmt.Errorf("error: Nested function '%v' evaluation failed\n%s", arg, err)
			}
			arguments[ind] = nestresult
		}
	}

	if extendedFunctions[funcName] != nil {
		return extendedFunctions[funcName].Call(arguments)
	}

	switch funcName {
	case "+":
		return plus(arguments)
	case "concat":
		return strconcat(arguments)
	}
	return nil, fmt.Errorf("error: Function '%s' is not defined", funcName)
}
