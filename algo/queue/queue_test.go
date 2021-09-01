package queue

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {

}

func TestQueue_EnQueue(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue("meow")
	q.EnQueue("2")
	if q.front != 3 {
		t.Errorf("front is not 3 it's %d", q.front)
	}
	q.DeQueue()
	if q.rear != 1 {
		t.Errorf("rear is not 3 it's %d", q.rear)
	}
	q.EnQueue(2)
	q.EnQueue(3)
	for {
		v := q.DeQueue()
		if v == nil {
			break
		}
		fmt.Println(v)
	}
	if q.front != 0 {
		t.Errorf("front is not 0 it's %d", q.front)
	}
	if q.rear != 0 {
		t.Errorf("rear is not 0 it's %d", q.rear)
	}
}

func TestQueue_IsFull(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	if q.front != 0 {
		t.Errorf("front is not 0 it's %d", q.front)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue("meow")
	q.EnQueue("2")
	q.EnQueue(2)
	q.EnQueue(3)

	for {
		v := q.DeQueue()
		if v == nil {
			break
		}
		fmt.Println(v)
	}
	if !q.IsEmpty() {
		t.Errorf("queue is not empty with front: %d, rear: %d", q.front, q.rear)
	}
}
