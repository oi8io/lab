package main

import "testing"
import "fmt"

func TestSlice(t *testing.T)  {
	var array [10]int
    var slice = array[5:6]
    fmt.Println("lenth of slice: ", len(slice))
    fmt.Println("capacity of slice: ", cap(slice))
    fmt.Println(&slice[0] == &array[5])
}
func TestSlice2(t *testing.T)  {
    var slice []int
    fmt.Println("capacity of slice: ", cap(slice))
    slice = append(slice, 1)
    fmt.Println("capacity of slice: ", cap(slice))
    slice = append(slice, 2)
    fmt.Println("capacity of slice: ", cap(slice))
    slice = append(slice,  3)
    fmt.Println("capacity of slice: ", cap(slice))
    newSlice := AddElement(slice, 4)
    fmt.Println("capacity of slice: ", cap(slice))
    fmt.Println("newSlice is same as slice? ",&slice[0] == &newSlice[0])
}


func TestSlice3(t *testing.T)  {
    orderLen := 5
    order := make([]uint16, 2 * orderLen)
    pollorder := order[:orderLen:orderLen]
    lockorder := order[orderLen:][:orderLen:orderLen]
    fmt.Println("len(pollorder) = ", len(pollorder))
    fmt.Println("cap(pollorder) = ", cap(pollorder))
    fmt.Println("len(lockorder) = ", len(lockorder))
    fmt.Println("cap(lockorder) = ", cap(lockorder))
}