package tests

import (
	"TAREA1/estructuras"
	"testing"
)

// TEST DE PUSH Y PEEK
func TestStackPushPeek(t *testing.T) {

	stack := estructuras.NewStack()

	stack.Push(5)
	stack.Push(10)

	top, ok := stack.Peek()

	if !ok {
		t.Error("Peek returned false but stack is not empty")
	}

	if top != 10 {
		t.Errorf("Expected 10 but got %d", top)
	}

}

// TEST DE POP
func TestStackPop(t *testing.T) {

	stack := estructuras.NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, ok := stack.Pop()

	if !ok {
		t.Error("Pop failed")
	}

	if val != 3 {
		t.Errorf("Expected 3 but got %d", val)
	}

}

// TEST DE ISEMPTY
func TestStackIsEmpty(t *testing.T) {

	stack := estructuras.NewStack()

	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}

}

// TEST DE POP EN STACK VACÍO
func TestPopEmptyStack(t *testing.T) {

	stack := estructuras.NewStack()

	_, ok := stack.Pop()

	if ok {
		t.Error("Pop should fail on empty stack")
	}

}
