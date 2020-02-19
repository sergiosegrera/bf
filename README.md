# BF
[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter written in Golang

## Usage
### Command line
```
$ brainfuck [FILE]
```
### In code
```
// Create interpreter
interpreter = NewInterpreter()
// Run Code
interpreter.Run(">+>+<-<-")
```
