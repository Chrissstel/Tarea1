package estructuras

type Queue struct {
	items []int
}

// Constructor
func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

// Enqueue
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// DEQUEUE
// regresa el primer elemento y si se pudo o no (en caso de estar vacío)
func (q *Queue) Dequeue() (int, bool) {

	if len(q.items) == 0 {
		return 0, false
	}

	element := q.items[0]

	q.items = q.items[1:]

	return element, true
}

// Peek
// regresa el elemento y si se pudo o no (en caso de estar vacío)
func (q *Queue) Peek() (int, bool) {

	if len(q.items) == 0 {
		return 0, false
	}

	return q.items[0], true
}

// IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
