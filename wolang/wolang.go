package wolang

import (
	"fmt"
)

func Eval(expr interface{}) (interface{}, error) {
	if res, err := eval(expr); err == nil {
		return res.GetValue(), nil
	} else {
		return nil, err
	}
}

func eval(expr interface{}) (result DataType, err error) {
	switch v := expr.(type) {
	case String:
		return v, nil
	case Float:
		return v, nil
	case Integer:
		return v, nil
	case Array:
		return evalFCall(v.value)
	}
	return nil, fmt.Errorf("error: Illegal expression type %T", expr)
}

func evalFCall(expr []DataType) (DataType, error) {
	funcName, isFName := expr[0].GetValue().(string)
	if !isFName {
		return nil, fmt.Errorf("error: Illegal call. %v is not a function name", expr[0].GetValue())
	}
	arguments := expr[1:]

	for ind, arg := range arguments {
		if argarray, ok := arg.(Array); ok {
			nestresult, err := evalFCall(argarray.value)
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
