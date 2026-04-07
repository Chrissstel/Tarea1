package tests

import (
	"TAREA1/estructuras"
	"testing"
)

// Test de enqueue y peek
func TestQueueEnqueuePeek(t *testing.T) {

	queue := estructuras.NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)

	val, ok := queue.Peek()

	if !ok {
		t.Error("Peek failed")
	}

	if val != 1 {
		t.Errorf("Expected 1 but got %d", val)
	}

}

// test de FIFO
func TestQueueDequeue(t *testing.T) {

	queue := estructuras.NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, _ := queue.Dequeue()

	if val != 1 {
		t.Errorf("Expected 1 but got %d", val)
	}

}

// test de queue vacía
func TestQueueEmpty(t *testing.T) {

	queue := estructuras.NewQueue()

	if !queue.IsEmpty() {
		t.Error("Queue should be empty")
	}

	queue.Enqueue(5)

	if queue.IsEmpty() {
		t.Error("Queue should not be empty")
	}

}
