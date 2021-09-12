package Repository

import (
	"errors"
	"sync"
)

type inMemoryDB struct {
	m    map[string]string
	lock sync.RWMutex
}

var KeyNotFoundErr = errors.New("this key does not exist")
var inMemoryRepositoryInstance *inMemoryDB = nil

// Initializes an empty in memory repository instance
func GetInMemoryRepositoryInstance() *inMemoryDB {
	if inMemoryRepositoryInstance == nil {
		inMemoryRepositoryInstance = &inMemoryDB{m: make(map[string]string)}
	}
	return inMemoryRepositoryInstance
}

// Get value based on given key
func (db *inMemoryDB) Get(key string) (*string, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	v, ok := db.m[key]
	if !ok {
		return nil, KeyNotFoundErr
	}
	return &v, nil
}

// Set given key value pair
func (db *inMemoryDB) Set(key, val string) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.m[key] = val
	return nil
}
