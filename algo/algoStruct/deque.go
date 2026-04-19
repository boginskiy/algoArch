package main

import "fmt"

// Deque - "double ended queue" - «очередь с двумя концами»
// В основе структура данных список (list) на кольцевом буфере.

type DeQueueOfRingBuff struct {
	queue []int
	size  int
	max   int
	Head  int
	Tail  int
}

func NewDeQueueOfRingBuff(n int) *DeQueueOfRingBuff {
	return &DeQueueOfRingBuff{
		queue: make([]int, n),
		max:   n,
	}
}

func (d *DeQueueOfRingBuff) IsEmpty() bool {
	return d.size == 0
}

func (d *DeQueueOfRingBuff) Pprint() {
	fmt.Println(d.queue)
}

func (d *DeQueueOfRingBuff) PushHead(elem int) {
	if d.size != d.max {
		d.queue[d.Head] = elem
		d.Head = (d.Head + 1) % d.max
		d.size++
	}
}

func (d *DeQueueOfRingBuff) PopHead() (elem int) {
	if !d.IsEmpty() {
		d.Head = (d.Head - 1) % d.max

		if d.Head < 0 {
			d.Head = d.max + d.Head
		}

		elem = d.queue[d.Head]
		d.queue[d.Head] = 0
		d.size--
	}
	return elem
}

func (d *DeQueueOfRingBuff) PushTail(elem int) {
	if d.size != d.max {
		d.Tail = (d.Tail - 1) % d.max

		if d.Tail < 0 {
			d.Tail = d.max + d.Tail
		}

		d.queue[d.Tail] = elem
		d.size++
	}
}

func (d *DeQueueOfRingBuff) PopTail() (elem int) {
	if !d.IsEmpty() {
		elem = d.queue[d.Tail]
		d.queue[d.Tail] = 0
		d.Tail = (d.Tail + 1) % d.max
		d.size--
	}
	return elem
}

func main() {
	NewDeQueueOfRingBuff(4)

}
