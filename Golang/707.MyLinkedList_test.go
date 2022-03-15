package Golang

import "testing"

func TestMyLinkedList(t *testing.T) {
	obj := ConstructorLinkedList()
	obj.AddAtTail(1)
	obj.AddAtTail(3)
	obj.AddAtTail(5)
	obj.AddAtTail(7)
	obj.AddAtTail(9)
	param_1 := obj.Get(2)
	t.Logf("Get:%v", param_1)
	obj.AddAtHead(0)
	obj.AddAtTail(11)
	obj.AddAtIndex(2, 4)
	obj.DeleteAtIndex(4)
}
