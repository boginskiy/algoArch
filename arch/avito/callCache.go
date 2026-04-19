package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

func hasher(s string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return int(h.Sum32())
}

// Реализация InMemory-cache

type Cache interface {
	Set(k, v string)
	Get(k string) (string, bool)
}

type InMemoryCache struct {
	shards []Shard
}

func NewInMemoryCache(num int) *InMemoryCache {
	shards := make([]Shard, 0, num)

	for i := 0; i < num; i++ {
		shards = append(shards, Shard{data: make(map[string]string)})
	}
	return &InMemoryCache{shards: shards}
}

func (i *InMemoryCache) Set(k, v string) {
	shardID := hasher(k) % len(i.shards)

	// Print
	fmt.Println(shardID)

	i.shards[shardID].Set(k, v)
}

func (i *InMemoryCache) Get(k string) (string, bool) {
	shardID := hasher(k) % len(i.shards)
	return i.shards[shardID].Get(k)
}

type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}

func (i *Shard) Set(k, v string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.data[k] = v
}

func (i *Shard) Get(k string) (string, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	v, ok := i.data[k]
	return v, ok
}

func main() {
	cache := NewInMemoryCache(5)

	cache.Set("A", "100")
	cache.Set("B", "200")
	cache.Set("C", "300")

}
