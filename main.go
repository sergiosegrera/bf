package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	interpreter := NewInterpreter()
	if len(os.Args) > 1 {
		code, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		interpreter.Run(string(code))
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">>> ")
			code, _ := reader.ReadString('\n')
			interpreter.Run(code)
		}
	}
}
