package link

import (
	"fmt"
	"testing"
)

var list *Array

func init() {
	list = NewArray()
}

func TestArray_Insert(t *testing.T) {
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()

	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()

	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()

	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()
}
func TestArray_Print(t *testing.T) {
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
	return
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()

	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()

	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.BulkInsert(0, []interface{}{1, 2, 3, 4, 5})
	list.Print()
}
