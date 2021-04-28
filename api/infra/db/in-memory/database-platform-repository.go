package inmemory

import (
	"errors"
	"sync"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
)

const (
	itemsPerPage = 10
)

type DatabasePlatformLocalStorageRepository struct {
	platforms protocols.PlatformType
	mutex     *sync.Mutex
}

func NewDatabasePlatformLocalStorageRepository(imoveis protocols.PlatformType) *DatabasePlatformLocalStorageRepository {
	return &DatabasePlatformLocalStorageRepository{
		platforms: imoveis,
		mutex:     new(sync.Mutex),
	}
}

func (d *DatabasePlatformLocalStorageRepository) GetProperties(platform string, page int) (protocols.ReturnPlatformResult, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	var response protocols.ReturnPlatformResult

	if platform == "zap" {
		res := d.platforms.Zap
		paginated := paginate(res, page)
		response = paginated
	} else if platform == "vivareal" {
		res := d.platforms.VivaReal
		paginated := paginate(res, page)

		response = paginated
	} else {
		return response, errors.New("Not found platform")
	}

	return response, nil
}

func paginate(data []protocols.Imovel, page int) protocols.ReturnPlatformResult {
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage

	if start > len(data) {
		return protocols.ReturnPlatformResult{}
	}

	if stop > len(data) {
		stop = len(data)
	}

	response := protocols.ReturnPlatformResult{
		PageNumber: page,
		PageSize:   itemsPerPage,
		TotalCount: len(data),
		Listings:   data[start:stop],
	}

	return response
}
