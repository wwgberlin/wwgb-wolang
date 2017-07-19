package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/julianbachiller/wwgb-wolang/wolang"
)

const (
	ERR_METHOD_UNDEFINED = "method not defined"
	ERR_ARGS_TOO_MANY = "too many arguments"
	ERR_ARGS_TOO_FEW = "too few arguments"
	ERR_ARGS_INVALID = "arguments are invalid"
)

func largerThan(terms []DataType) (result DataType, err error) {
	return nil, fmt.Errorf(ERR_METHOD_UNDEFINED)
}

func smallerThan(terms []DataType) (result DataType, err error) {
	return nil, fmt.Errorf(ERR_METHOD_UNDEFINED)
}

func equals(terms []DataType) (result DataType, err error) {
	return nil, fmt.Errorf(ERR_METHOD_UNDEFINED)
}

func registerExtensions() {
	var extensions []ExtFuncDef = []ExtFuncDef{
		{">", largerThan},
		{"<", smallerThan},
		{"==", equals},
	}

	for _, ex := range extensions {
		RegExtFunc(ex)
	}
}

func main() {

	for {
		run()
	}
}

func run() {

	registerExtensions();

	var result interface{}
	var parsed DataType
	var unparsed string

	var input string
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("wolang> ")

	if input, err = reader.ReadString('\n'); err != nil {
		panic(err.Error())
	}

	// Remove trailing newline chars \n and \r for Win. compat
	unparsed = strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")

	// Parse and eval until no more input
	for len(unparsed) > 0 {
		unparsed, parsed, err = Parse(unparsed)
		if err != nil {
			// Display error and break to main eval loop
			fmt.Println(err.Error())
			break
		}

		result, err = Eval(parsed)
		if err != nil {
			// Display error and break to main eval loop
			fmt.Println(err.Error())
			break
		}

		// Display result
		fmt.Println(result)
	}
}
