package main

import (
	"fmt"

	"github.com/julianbachiller/wwgb-wolang/wolang"
)

func main() {

	//
	// concatenate strings
	//

	expr1 := []interface{}{"concat", "a", "b", "cde"}
	result1, err1 := wolang.Eval(expr1)
	fmt.Println(result1, err1)

	//
	// add numbers
	//
	expr2 := []interface{}{"+", 1, 2, 7, 7}
	result2, err2 := wolang.Eval(expr2)
	fmt.Println(result2, err2)

	//
	// cause an error in concat
	//
	expr3 := []interface{}{"concat", "a", "b", 17}
	result3, err3 := wolang.Eval(expr3)
	fmt.Println(result3, err3)

	//
	// cause an error in add
	//
	expr4 := []interface{}{"+", 1, 2, true}
	result4, err4 := wolang.Eval(expr4)
	fmt.Println(result4, err4)

	// evaluate a single string
	fmt.Println(wolang.Eval("hi there"))

	// evaluate a single number
	fmt.Println(wolang.Eval(27))

	//
	// cause an error because of calling undefined function
	//
	expr5 := []interface{}{"minus", 5, 2}
	result5, err5 := wolang.Eval(expr5)
	fmt.Println(result5, err5)

	//
	// call an extended function
	//

	var always99 wolang.ExtFuncDef = wolang.ExtFuncDef{
		"always99",
		func(terms []interface{}) (result interface{}, err error) {
			return 99, nil
		},
	}
	wolang.RegExtFunc(always99)
	expr6 := []interface{}{"always99", 5, 2}
	result6, err6 := wolang.Eval(expr6)
	fmt.Println(result6, err6)

	//
	// evaluate a nested expression
	//

	nestedExpr := []interface{}{"+", 3, 4, []interface{}{"+", 5, 6}}
	result7, err7 := wolang.Eval(nestedExpr)
	fmt.Println(result7, err7)
}
