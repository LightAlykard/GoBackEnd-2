package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LightAlykard/GoBackEnd-2/HW-6/models"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	*redis.Client
	ttl time.Duration
}

func NewUserStorage(host, port string, ttl time.Duration) (UserStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // TODO: Add from config
		DB:       0,  // use default DB
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("try to ping to redis: %w", err)
	}
	c := &RedisClient{
		Client: client,
	}
	return c, nil
}
func (c *RedisClient) Close() error {
	return c.Client.Close()
}

func (c *RedisClient) GetUser(ctx context.Context, name string) (*models.User, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("GetUser done with context")
	default:
		data, err := c.Get(ctx, name).Bytes()
		if err == redis.Nil {
			// we got empty result, it's not an error
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		u := models.User{}
		err = json.Unmarshal(data, &u)
		if err != nil {
			return nil, fmt.Errorf("can't decode data: %s", err)
		}
		return &u, nil
	}
}
func (c *RedisClient) Create(ctx context.Context, user *models.User) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("user.Create done with context")
	default:
		err := c.Set(ctx, user.Name, user, c.ttl).Err()
		if err != nil {
			return fmt.Errorf("can't add data to redis: %s", err)
		}
		return nil
	}

}