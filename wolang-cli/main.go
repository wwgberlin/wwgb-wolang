package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/julianbachiller/wwgb-wolang/wolang"
)

func main() {

	var result interface{}
	var parsed interface{}

	var unparsed string

	var err error

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("wolang> ")
		if input, readErr := reader.ReadString('\n'); readErr != nil {
			panic(readErr.Error())
		} else {
			// Remove trailing newline chars \n and \r for Win. compat
			unparsed = strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")

			// Parse and eval until no more input
			for len(unparsed) > 0 {
				unparsed, parsed, err = wolang.Parse(unparsed)
				if err != nil {
					// Display error and break to main eval loop
					fmt.Println(err.Error())
					break
				}

				result, err = wolang.Eval(parsed)
				if err != nil {
					// Display error and break to main eval loop
					fmt.Println(err.Error())
					break
				}

				// Display result
				fmt.Println(result)
			}
		}
	}
}
