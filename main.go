package main

import (
	"TAREA1/estructuras"
	"fmt"
)

func main() {

	//IMPLEMENTACIÓN DE STACK
	stack := estructuras.NewStack()
	fmt.Println("is it empty?", stack.IsEmpty())

	stack.Push(5)
	stack.Push(7)
	stack.Push(10)

	top, _ := stack.Peek()
	fmt.Println("top element:", top)

	val, _ := stack.Pop()
	fmt.Println("Popped: ", val)
	val, _ = stack.Pop()
	fmt.Println("Popped: ", val)

	fmt.Println("is it empty?", stack.IsEmpty())
	top, _ = stack.Peek()
	fmt.Println("top element:", top)

	//IMPLEMENTACIÓN DE QUEUE
}
