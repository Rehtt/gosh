package utils

import (
	"sync"
	"time"
)

type Dog struct {
	timer *time.Timer
	times time.Duration
}

func NewWatchDog(w *sync.WaitGroup) *Dog {
	d := &Dog{times: 5 * time.Minute}
	t := time.NewTimer(d.times)
	d.timer = t
	go func() {
		select {
		case <-t.C:
			t.Stop()
			w.Done()
		}
	}()
	return d
}
func (d *Dog) FeedDog() {
	d.timer.Reset(d.times)
}
