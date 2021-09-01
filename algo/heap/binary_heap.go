package heap

import "fmt"

type Element int

type BinaryHeap struct {
	Capacity int
	Size     int
	Elements []Element
}

type PriorityQueue BinaryHeap

func NewPriorityQueue(capacity int) *PriorityQueue {
	h := new(PriorityQueue)
	h.Capacity = capacity
	h.Size = 0
	h.Elements = append(h.Elements, 0)
	return h
}
func BuildPriorityQueue(capacity int, eles []Element) *PriorityQueue {
	h := new(PriorityQueue)
	h.Capacity = capacity
	h.Elements = append(h.Elements, 0)
	h.Elements = append(h.Elements, eles...)
	h.Size = len(h.Elements)
	for i := h.Size / 2; i > 0; i-- {
		//fmt.Println(i)
		h.Down(i)
	}

	return h
}
func (h *PriorityQueue) Insert(element Element) {
	h.Elements = append(h.Elements, element)
	cur := h.Size
	fmt.Println("size ", h.Size)
	for {
		if h.Elements[cur] < h.Elements[cur/2] {
			fmt.Println("inserting ", h.Elements[cur], "<", h.Elements[cur/2])
			tmp := h.Elements[cur/2]
			h.Elements[cur/2] = h.Elements[cur]
			h.Elements[cur] = tmp
			cur = cur / 2
		} else {
			break
		}
		if cur == 1 {
			break
		}
	}
	h.Size = len(h.Elements)
}

func (h *PriorityQueue) Destroy() {

}

func (h *PriorityQueue) DeleteMin() {
	i := 1
	for {
		if 2*i+1 >= h.Size {
			fmt.Println(2*i+1, h.Size)
			break
		}
		if h.Elements[2*i] > h.Elements[2*i+1] {
			fmt.Println(h.Elements[2*i], ">", h.Elements[2*i+1])
			h.Elements[i] = h.Elements[2*i+1]
			i = 2*i + 1
		} else {
			fmt.Println(h.Elements[2*i], "<", h.Elements[2*i+1])
			h.Elements[i] = h.Elements[2*i]
			i = 2 * i
		}
	}
	h.Elements[i] = h.Elements[h.Size-1]
	h.Elements = h.Elements[:h.Size-1]
}

func (h *PriorityQueue) Down(i int) {
	for {
		if 2*i+1 >= h.Size {
			//fmt.Println("i:", i, 2*i+1, ">=", h.Size, "break")
			break
		}
		//fmt.Println(h.Elements[i], h.Elements[2*i], h.Elements[2*i+1])
		if h.Elements[2*i] < h.Elements[2*i+1] {
			if h.Elements[i] > h.Elements[2*i] {
				tmp := h.Elements[i]
				h.Elements[i] = h.Elements[2*i]
				h.Elements[2*i] = tmp
				i = i * 2
			} else {
				break
			}
		} else {
			if h.Elements[i] > h.Elements[2*i+1] {
				tmp := h.Elements[i]
				h.Elements[i] = h.Elements[2*i+1]
				h.Elements[2*i+1] = tmp
				i = i*2 + 1
			} else {
				break
			}
		}

	}
}

func (h *PriorityQueue) FindMin() {

}
