package util

import (
	"sync"
)

type LinkedList struct {
	pre   *LinkedList
	next  *LinkedList
	Value string
}
type Link struct {
	top    *LinkedList
	bottom *LinkedList
	size   int
	mu     sync.RWMutex
}

func NewLink() *Link {
	return new(Link)
}
func (l *Link) SetMaxSize(size int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.top == nil {
		l.top = new(LinkedList)
		l.bottom = l.top
	}
	index := l.top
	for i := 1; i < size; i++ {
		if index.next == nil {
			index.next = new(LinkedList)
			index.next.pre = index
		}
		index = index.next
	}
	index.next = l.top
	l.top.pre = index
}
func (l *Link) Write(value string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.top == l.bottom && l.top.Value != "" {
		l.top = l.top.next
		l.size--
	}
	l.bottom.Value = value
	l.bottom = l.bottom.next
	l.size++
}
func (l *Link) Range() (out []string) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	index := l.top
	for i := 0; i < l.size; i++ {
		out = append(out, index.Value)
		index = index.next
	}

	return out
}