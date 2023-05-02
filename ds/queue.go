package ds

type Queue struct {
	capacity int
	values   []interface{}
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		capacity: capacity,
		values:   make([]interface{}, 0),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Size() <= 0
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) Enqueueing(value interface{}) error {
	if q.Size()+1 > q.capacity {
		return ErrCapacityLimit
	}
	copiedValue := make([]interface{}, q.Size()+1)
	copy(copiedValue[1:], q.values)

	q.values = copiedValue
	q.values[0] = value
	return nil
}

func (q *Queue) Dequeueing() interface{} {
	if q.Size() <= 0 {
		return ErrNoSuchElement
	}
	value := q.Peek()
	q.values = q.values[:q.GetLastIndex()]
	return value
}

func (q *Queue) Peek() interface{} {
	if q.Size() <= 0 {
		return ErrNoSuchElement
	}
	return q.values[q.GetLastIndex()]
}

func (q *Queue) GetLastIndex() int {
	return q.Size() - 1
}

func (q *Queue) Clear() error {
	q.values = make([]interface{}, 0)
	return nil
}
