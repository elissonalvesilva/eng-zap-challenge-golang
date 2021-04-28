package inmemory

import (
	"errors"
	"sync"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"
)

type DatabaseLocalStorage struct {
	platforms map[string][]entity.Imovel
	mutex     *sync.Mutex
}

func NewDatabaseLocalStorage() *DatabaseLocalStorage {
	return &DatabaseLocalStorage{
		platforms: make(map[string][]entity.Imovel),
		mutex:     new(sync.Mutex),
	}
}

func (d *DatabaseLocalStorage) GetProperties(platform string) ([]entity.Imovel, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	response := d.platforms[platform]

	if len(response) == 0 {
		return nil, errors.New("Not found platform")
	}

	return response, nil
}
