package cache

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	dk *bolt.DB
)

type disk struct{}

func (d *disk) Init() {
	config := &bolt.Options{Timeout: 1 * time.Second}
	db, err := bolt.Open("disk.db", 0600, config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Default"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	dk = db
}

func (d *disk) Set(key string, value string, ttl time.Duration) error {
	err := dk.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Default"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

func (d *disk) Get(key string, defaultValue string) (value string, err error) {
	err = dk.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Default"))
		value = string(b.Get([]byte(key)))
		return nil
	})
	return
}

func (d *disk) Remove(key string) error {
	err := dk.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Default"))
		err := b.Delete([]byte(key))
		return err
	})
	return err
}

func (d *disk) Exists(key string) (exists bool) {
	_ = dk.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Default"))
		value := string(b.Get([]byte(key)))
		exists = value != ""
		return nil
	})
	return
}

func Disk() *disk {
	d := &disk{}
	d.Init()
	return d
}
