package utils

import (
	"sync"
)

func ParallelFor(begin, end, lim int, fn func(i int)) {
	var wg sync.WaitGroup
	wg.Add(end - begin)
	limits := make(chan struct{}, lim)
	for i := begin; i < end; i++ {
		limits <- struct{}{}
		go func(i int) {
			fn(i)
			<-limits
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func ParallelForEnd(begin, end, lim int, fn func(i int) bool) {
	var wg sync.WaitGroup
	wg.Add(end - begin)
	limits := make(chan struct{}, lim)
	isEnd := false
	for i := begin; i < end; i++ {
		limits <- struct{}{}
		go func(i int) {
			if !isEnd {
				if !fn(i) {
					isEnd = true
				}
			}
			<-limits
			wg.Done()
		}(i)
	}
	wg.Wait()
}
