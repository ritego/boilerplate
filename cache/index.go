package cache

import (
	"time"

	"github.com/aellacredit/jara/config"
)

type CacheContract interface {
	Init()
	Set(key string, value string, ttl time.Duration) error
	Get(key string, defaultValue string) (value string, err error)
	Remove(key string) error
	Exists(key string) (exists bool)
}

const (
	MEMORY_CACHE_DRIVER = "memory"
	DISK_CACHE_DRIVER   = "disk"
)

func Default() CacheContract {
	var d CacheContract

	if config.String("CACHE_DRIVER") == DISK_CACHE_DRIVER {
		d = &disk{}
	} else {
		d = &memory{}
	}

	return d
}
