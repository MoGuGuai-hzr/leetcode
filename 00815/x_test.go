package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(1)
	l.PushBack(2)
	e := l.Back()
	l.Remove(e)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i)
	}
}
