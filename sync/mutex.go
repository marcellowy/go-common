package sync

import "time"

type Mutex interface {
	Lock(key string, expires time.Duration) error
	Unlock() error
}
