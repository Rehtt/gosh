package util

import (
	"fmt"
	"testing"
)

func TestNewLink(t *testing.T) {
	l := NewLink()
	l.SetMaxSize(4)
	l.InsertValue("1")
	l.InsertValue("2")
	//l.InsertValue("3")
	//l.InsertValue("4")
	//l.InsertValue("5")
	fmt.Println(l.Range())
}
