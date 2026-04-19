package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Адрес и порт Redis-сервера
		Password: "",               // Пароль (если есть)
		DB:       0,                // Индекс базы данных (по умолчанию 0)
	})

	// Тестовое подключение
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("failed to connect to Redis: %w", err))
	}
	fmt.Println("Ping response from Redis:", pong)

	// Остальные операции с Redis
}
