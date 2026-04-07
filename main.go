package main

import (
	"TAREA1/estructuras"
	"fmt"
)

func main() {

	//IMPLEMENTACIÓN DE STACK
	stack := estructuras.NewStack()
	fmt.Println("\nStack demo")
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
	queue := estructuras.NewQueue()
	fmt.Println("\nQueue demo")

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	front, _ := queue.Peek()
	fmt.Println("Front:", front)
	val, _ = queue.Dequeue()
	fmt.Println("Dequeued:", val)
	val, _ = queue.Dequeue()
	fmt.Println("Dequeued:", val)

	fmt.Println("Is empty?", queue.IsEmpty())

	//IMPLEMENETACIÓN DICTIONARY
}
