package estructuras

type Stack struct {
	items []int
}

// CONSTRUCTOR
func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

// PUSH
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// POP
// regresa el valor y si se pudo o no (en caso de que esté vacío)
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	index := len(s.items) - 1
	element := s.items[index]

	s.items = s.items[:index]

	return element, true
}

// PEEK
// regresa el último elemento y si se pudo o no (en caso de que esté vacío)
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// ISEMPTY
// regresa si está vacío o no
func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
