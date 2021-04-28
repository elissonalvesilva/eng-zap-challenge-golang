package inmemory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"
)

type DatabasePlatformLocalStorageRepository struct {
	platforms map[string][]entity.Imovel
	mutex     *sync.Mutex
}

func NewDatabasePlatformLocalStorageRepository() *DatabasePlatformLocalStorageRepository {
	return &DatabasePlatformLocalStorageRepository{
		platforms: make(map[string][]entity.Imovel),
		mutex:     new(sync.Mutex),
	}
}

func (d *DatabasePlatformLocalStorageRepository) GetProperties(platform string) ([]entity.Imovel, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	fmt.Println(d.platforms)
	response := d.platforms[platform]

	if len(response) == 0 {
		return nil, errors.New("Not found platform")
	}

	return response, nil
}
