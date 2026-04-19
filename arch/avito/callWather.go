package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func WatherForecast(city string) int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

type Data struct {
	Temperatures map[string]int
	mu           sync.RWMutex
}

func NewData(ctx context.Context, interval time.Duration) *Data {
	ticker := time.NewTicker(interval)
	data := &Data{}

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				data.UpdateTemperature(ctx)
			}
		}
	}()

	return data
}

func (d *Data) GetTemperature(city string) (int, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	v, ok := d.Temperatures[city]

	if !ok {
		return 0, fmt.Errorf("city %s not found", city)
	}
	return v, nil
}

// Попробуй модернизировать и проверить работу.
func (d *Data) UpdateTemperature(ctx context.Context) {
	wg := &sync.WaitGroup{}

	for city := range d.Temperatures {
		wg.Add(1)

		go func(city string) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				tmp := WatherForecast(city)
				d.mu.Lock()
				d.Temperatures[city] = tmp
				d.mu.Unlock()
			}

		}(city)
	}

	wg.Wait()
}

func main() {
	//
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	data := NewData(ctx, 1*time.Minute)

	http.HandleFunc("/wather", func(w http.ResponseWriter, r *http.Request) {
		temp, err := data.GetTemperature("Moscow")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		fmt.Fprintf(w, "{\"temperature\":%d}\n", temp)
	})

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
