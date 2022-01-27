package cache

import (
	"log"
	"time"

	"github.com/tidwall/buntdb"
)

var (
	mem *buntdb.DB
)

type memory struct{}

func (m *memory) Init() {
	db, err := buntdb.Open("memory.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config := buntdb.Config{
		SyncPolicy:           buntdb.EverySecond,
		AutoShrinkPercentage: 100,
		AutoShrinkMinSize:    32,
		AutoShrinkDisabled:   false,
	}
	if err := db.SetConfig(config); err != nil {
		log.Fatal(err)
	}

	mem = db
}

func (m *memory) Set(key string, value string, ttl time.Duration) error {
	err := mem.Update(func(tx *buntdb.Tx) error {
		var config *buntdb.SetOptions

		if ttl == 0 {
			config = nil
		} else {
			config = &buntdb.SetOptions{Expires: true, TTL: ttl}
		}

		_, _, err := tx.Set(key, value, config)
		return err
	})
	return err
}

func (m *memory) Get(key string, defaultValue string) (value string, err error) {
	err = mem.View(func(tx *buntdb.Tx) error {
		value, err = tx.Get(key)
		if err != nil {
			return err
		}

		if value == "" {
			value = defaultValue
		}

		return err
	})
	return
}

func (m *memory) Remove(key string) error {
	err := mem.View(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *memory) Exists(key string) (exists bool) {
	_ = mem.View(func(tx *buntdb.Tx) error {
		_, err := tx.Get(key)
		exists = err != nil
		return err
	})
	return
}

func Memory() *memory {
	m := &memory{}
	m.Init()
	return m
}
