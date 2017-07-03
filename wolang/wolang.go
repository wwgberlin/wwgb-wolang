package main

import (
	"fmt"
)

func eval(expr interface{}) (result interface{}, err error) {
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
	funcName := expr[0].(string)
	arguments := expr[1:]

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

func main() {

	//
	// concatenate strings
	//

	expr1 := []interface{}{"concat", "a", "b", "cde"}
	result1, err1 := eval(expr1)
	fmt.Println(result1, err1)

	//
	// add numbers
	//
	expr2 := []interface{}{"+", 1, 2, 7, 7}
	result2, err2 := eval(expr2)
	fmt.Println(result2, err2)

	//
	// cause an error in concat
	//
	expr3 := []interface{}{"concat", "a", "b", 17}
	result3, err3 := eval(expr3)
	fmt.Println(result3, err3)

	//
	// cause an error in add
	//
	expr4 := []interface{}{"+", 1, 2, true}
	result4, err4 := eval(expr4)
	fmt.Println(result4, err4)

	// evaluate a single string
	fmt.Println(eval("hi there"))

	// evaluate a single number
	fmt.Println(eval(27))

	//
	// cause an error because of calling undefined function
	//
	expr5 := []interface{}{"minus", 5, 2}
	result5, err5 := eval(expr5)
	fmt.Println(result5, err5)

	//
	// call an extended function
	//
	expr6 := []interface{}{"always99", 5, 2}
	result6, err6 := eval(expr6)
	fmt.Println(result6, err6)
}
