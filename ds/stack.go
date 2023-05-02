package ds

type Stack struct {
	values []interface{}
}

func NewStack() *Stack {
	return &Stack{
		values: make([]interface{}, 0),
	}
}

func (s Stack) IsEmpty() bool {
	return s.Size() <= 0
}

func (s Stack) Size() int {
	return len(s.values)
}

func (s *Stack) Push(val interface{}) error {
	s.values = append(s.values, val)
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Size() <= 0 {
		return ErrNoSuchElement
	}
	return s.values[s.GetLastIndex()]
}

func (s *Stack) Pop() interface{} {
	value := s.values[s.GetLastIndex()]
	s.values = s.values[:s.GetLastIndex()]
	return value
}

func (s *Stack) GetLastIndex() int {
	return s.Size() - 1
}
