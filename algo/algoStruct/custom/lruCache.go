package main

import "container/list"

// Пример задачи: Реализация LRU-кэша (часто встречается в VK)
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.([2]int)[1]
	}
	return -1
}
