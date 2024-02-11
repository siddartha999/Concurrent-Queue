package main

import (
	"fmt"
	"sync"
)

var wgE sync.WaitGroup
var wgD sync.WaitGroup

type ConcurrentQueue struct {
	queue []int
	mutex sync.Mutex
}

func (q* ConcurrentQueue) Enqueue(item int) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, item)
}

func (q* ConcurrentQueue) Dequeue() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.queue) == 0 {
		panic("Cannot remove from an empty Queue")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q* ConcurrentQueue) Size() int {
	return len(q.queue)
}

func main() {
	qq := ConcurrentQueue{
		queue: make([]int, 0),
	}

	for i := 0; i < 10000; i++ {
		wgE.Add(1)
		go func () {
			qq.Enqueue(i)
			wgE.Done()
		}()
	}
	wgE.Wait()
	fmt.Println("Thread Safe Concurrent Queue, Enqueue operation with 10K inserts: %s", qq.Size())

	for i := 0; i < 10000; i++ {
		wgD.Add(1)
		go func() {
			qq.Dequeue()
			wgD.Done()
		}()
	}
	wgD.Wait()
	fmt.Println("Thread Safe Concurrent Queue, Dequeue operation after 10K removals: %s", qq.Size())
}