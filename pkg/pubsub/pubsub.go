package pubsub

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	Redis *redis.Client
}

func New(address string, password string) *Service {
	redis := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	err := redis.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("cannot ping redis: %v", err)
	}

	return &Service{
		Redis: redis,
	}
}

func (s *Service) Publish(ctx context.Context, channel string, value interface{}) error {
	err := s.Redis.Publish(ctx, channel, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Subscribe(ctx context.Context, channel string, msg chan []byte) error {
	// go-redis automatically reconnects on error
	subscriber := s.Redis.Subscribe(ctx, channel)
	ch := subscriber.Channel()

	go func() {
		for resp := range ch {
			msg <- []byte(resp.Payload)
		}
	}()
	return nil
}
