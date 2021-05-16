package redis

import (
	"context"
	"fmt"
	"time"

	v8 "github.com/go-redis/redis/v8"
)

type Mutex struct {
	key     string
	client  *v8.Client
	Timeout time.Duration
}

// NewMutex 新建
func NewMutex(client *v8.Client) *Mutex {
	return &Mutex{client: client, Timeout: time.Second * 3}
}

// Lock 上锁
func (m *Mutex) Lock(key string, expires time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.Timeout)
	defer cancel()
	m.key = key
	cmd := m.client.SetNX(ctx, key, "1", expires)
	if cmd.Err() == v8.Nil {
		return nil
	} else if cmd.Err() != nil {
		return cmd.Err()
	}
	return fmt.Errorf("key %s exists", key)
}

// Unlock 释放
func (m *Mutex) Unlock() error {
	ctx, cancel := context.WithTimeout(context.Background(), m.Timeout)
	defer cancel()
	return m.client.Del(ctx, m.key).Err()
}
