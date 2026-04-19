package main

// Queue - В основе структура данных список (list) на кольцевом буфере.

type QueueOfRingBuff struct {
	queue []int
	max   int // максимально возможное количество элементов в очереди
	head  int // индекс, по которому нужно извлекать элемент, если очередь не пустая
	tail  int // индекс, по которому нужно добавлять элемент, если в очереди есть место
	size  int // размер очереди
}

func NewQueueOfRingBuff(n int) *QueueOfRingBuff {
	return &QueueOfRingBuff{
		queue: make([]int, n),
		max:   n,
	}
}

func (q *QueueOfRingBuff) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueOfRingBuff) Push(elem int) {
	if q.size != q.max {
		q.queue[q.tail] = elem
		q.tail = (q.tail + 1) % q.max
		q.size++
	}
}

func (q *QueueOfRingBuff) Pop() int {
	if q.IsEmpty() {
		return 0
	}
	elem := q.queue[q.head]
	q.queue[q.head] = 0
	q.head = (q.head + 1) % q.max
	q.size--
	return elem
}
