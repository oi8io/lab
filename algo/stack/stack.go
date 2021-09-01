package stack

type Element int
type Stack struct {
	value Element
	next  *Stack
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.next == nil
}

func (s *Stack) Dispose() {

}

func (s *Stack) Push(element Element) *Stack {
	i := &Stack{value: element,}
	if s.IsEmpty() {
		s.next = i
	} else {
		i.next = s.next
		s.next = i
	}
	return s
}

func (s *Stack) Pop() (*Stack, Element) {
	ele := s.next
	if s.IsEmpty() {
		return s, 0
	}
	s.next = ele.next
	return s, ele.value
}
