package pokecache

import "time"

type Cache struct {
	createdAt time.Time
	val       []byte
}

func Add(key string, val []byte) {

}

func Get(key string) ([]byte, bool) {}

func reapLoop()
