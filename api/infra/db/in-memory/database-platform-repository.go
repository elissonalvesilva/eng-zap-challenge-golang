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
	Platforms protocols.PlatformType
	Mutex     *sync.Mutex
}

func NewDatabasePlatformLocalStorageRepository(imoveis protocols.PlatformType) *DatabasePlatformLocalStorageRepository {
	return &DatabasePlatformLocalStorageRepository{
		Platforms: imoveis,
		Mutex:     new(sync.Mutex),
	}
}

func (d *DatabasePlatformLocalStorageRepository) GetProperties(platform string, page int) (protocols.ReturnPlatformResult, protocols.ErrorResponse) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	var response protocols.ReturnPlatformResult
	var err = protocols.ErrorResponse{}
	if platform == "zap" {
		res := d.Platforms.Zap
		paginated := paginate(res, page)
		response = paginated
	} else if platform == "vivareal" {
		res := d.Platforms.VivaReal
		paginated := paginate(res, page)

		response = paginated
	} else {
		err = protocols.ErrorResponse{
			Message: errors.New("Not found platform").Error(),
		}
		return response, err
	}

	return response, err
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
