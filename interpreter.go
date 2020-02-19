package main

import (
	"bufio"
	"fmt"
	"os"
)

// Brainfuck keywords
const (
	GREATER  = 62
	LESSER   = 60
	PLUS     = 43
	MINUS    = 45
	DOT      = 46
	COMMA    = 44
	LBRACKET = 91
	RBRACKET = 93
)

type Interpreter struct {
	memory   *MemoryManager
	pointer  int
	brackets [][]int
	reader   *bufio.Reader
}

func NewInterpreter() *Interpreter {
	return &Interpreter{NewMemory(), 0, make([][]int, 0, 10), bufio.NewReader(os.Stdin)}
}

func (in Interpreter) Run(code string) {
	// Parse through code
	i := 0
	for i < len(code) {
		switch code[i] {
		case GREATER:
			in.pointer++
		case LESSER:
			// Pointer can't be negative
			if in.pointer > 0 {
				in.pointer--
			}
		case PLUS:
			in.memory.Add(in.pointer)
		case MINUS:
			in.memory.Substract(in.pointer)
		case DOT:
			fmt.Print(string(in.memory.Get(in.pointer)))
		case COMMA:
			// Reads only first rune of string
			fmt.Print("> ")
			input, _ := in.reader.ReadString('\n')
			in.memory.Put(i, int(input[0]))
		case LBRACKET:
			// If "[" is found insert a slice into c.brackets that contains the index where the "[" was found and the current location of the pointer
			in.brackets = append(in.brackets, []int{i, in.pointer})
		case RBRACKET:
			// If "]" is found check if the corresponding pointer is 0. If not jump back the index to the corresponding "["
			if in.memory.Get(in.brackets[len(in.brackets)-1][1]) == 0 && len(in.brackets) > 0 {
				in.brackets = in.brackets[:len(in.brackets)-1]
			} else {
				i = in.brackets[len(in.brackets)-1][0]
			}
		}
		i++
	}
	fmt.Println("")
}
