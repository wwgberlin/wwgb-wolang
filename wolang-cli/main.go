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
	var input string
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("wolang> ")

	if input, err = reader.ReadString('\n'); err != nil {
		panic(err.Error())
	}

	_, parsed, err = Parse(strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if result, err = Eval(parsed); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Display result
	fmt.Println(result)
}
