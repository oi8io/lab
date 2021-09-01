package heap

import (
	"fmt"
	"testing"
)

var pg *PriorityQueue
func init() {
	//pg = NewPriorityQueue(8)
	//pg = BuildPriorityQueue(8,[]Element{13,21,16,24,31,19,68,65,26,32})
}

func TestPriorityQueue_Insert(t *testing.T) {
	fmt.Println(pg)
	pg.Insert(14)
	fmt.Println(pg)
}

func TestPriorityQueue_DeleteMin(t *testing.T) {
	pg = BuildPriorityQueue(8,[]Element{150,80,40,30,10,70,110,100,20,90,60,50,120,140,130})
	fmt.Println(pg)
}