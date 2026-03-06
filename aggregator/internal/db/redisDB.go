package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type StoreRS struct {
	RDB *redis.Client
	ctx context.Context
}

func NewStoreRS(ctx context.Context) *StoreRS {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return &StoreRS{
		RDB: rdb,
		ctx: ctx,
	}
}

func (s *StoreRS) Ping() bool {
	_, err := s.RDB.Ping(s.ctx).Result()
	if err != nil {
		return false
	}
	return true
}

func (s *StoreRS) Close() {
	s.RDB.Close()
}

func (s *StoreRS) GetDB() *redis.Client {
	return s.RDB
}
