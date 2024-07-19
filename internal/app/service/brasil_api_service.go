package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/brenddonanjos/multithreading_api/internal/entity"
)

type BrasilApiService struct{}

func NewBrasilApiService() *BrasilApiService {
	return &BrasilApiService{}
}

func (ba *BrasilApiService) FetchZipCode(zipCode string, startTime time.Time) (zipCodeInfo *entity.ZipCode, err error) {
	req, err := http.NewRequest(http.MethodGet, "https://brasilapi.com.br/api/cep/v1/"+zipCode, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read json response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var brasilApi entity.BrasilApi
	//Prepare and return struct
	err = json.Unmarshal(body, &brasilApi)
	if err != nil {
		return nil, err
	}

	if brasilApi.Cep == "" {
		return nil, errors.New("cep not found (brasilapi)")
	}

	executionTime := time.Since(startTime).String()

	zipCodeInfo = entity.NewZipCode(brasilApi.Cep, brasilApi.State, brasilApi.City, brasilApi.Neighborhood, brasilApi.Street, "Brasil API", executionTime)

	return zipCodeInfo, nil
}
