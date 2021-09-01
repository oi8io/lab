package queue

type QueueInf interface {
	Size() int
	IsEmpty() bool
	EnQueue(interface{})
	DeQueue() interface{}
}

const BuffSize = 100

type Queue struct {
	buff  *[BuffSize]interface{}
	rear  int
	front int
}

func NewQueue() *Queue {
	return &Queue{
		buff:  &[BuffSize]interface{}{},
		rear:  0,
		front: 0,
	}
}

func (q *Queue) Size() int {
	if q.front < q.rear {
		return q.front + BuffSize - q.rear
	}
	return q.front - q.rear
}

func (q *Queue) IsEmpty() bool {
	return q.rear == q.front
}

func (q *Queue) isFull() bool {
	return false
}

func (q *Queue) EnQueue(i interface{}) {
	q.buff[q.front] = i
	q.front += 1
	q.front = q.front % BuffSize
	return
}

func (q *Queue) DeQueue() interface{} {
	if q.buff[q.rear] == nil { //consume fail
		return nil
	}
	item := q.buff[q.rear]
	q.buff[q.rear] = nil
	q.rear += 1
	q.rear = q.rear % BuffSize
	return item
}
