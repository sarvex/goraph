package redblack

import (
	"fmt"
	"testing"
)

type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

func TestIsRed(t *testing.T) {
	if !isRed(NewNode(Int(5))) {
		t.Error("Expected Red")
	}
}

type Str string

// Less returns true if string(a) < string(b).
func (a Str) Less(b Interface) bool {
	return a < b.(Str)
}

func TestInsert1(t *testing.T) {
	root := NewNode(Str("S"))
	data := New(root)
	data.Insert(NewNode(Str("E")))
	data.Insert(NewNode(Str("A")))
	data.Insert(NewNode(Str("S")))
	nd := data.Search(Str("E"))
	if fmt.Sprintf("%v", nd.Key) != "E" {
		t.Errorf("Unexpected %v", nd.Key)
	}
	if fmt.Sprintf("%v", nd.Left.Key) != "A" {
		t.Errorf("Unexpected %v", nd.Left.Key)
	}
	if fmt.Sprintf("%v", nd.Right.Key) != "S" {
		t.Errorf("Unexpected %v", nd.Right.Key)
	}
	if !nd.Left.Black {
		t.Errorf("Left should be black but %v", nd.Left.Black)
	}
	if !nd.Right.Black {
		t.Errorf("Right should be black but %v", nd.Right.Black)
	}
}