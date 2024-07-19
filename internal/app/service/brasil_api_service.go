package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/brenddonanjos/multithreading_api/internal/entity"
)

type BrasilApiService struct {
	Ctx context.Context
}

func NewBrasilApiService(ctx context.Context) *BrasilApiService {
	return &BrasilApiService{
		Ctx: ctx,
	}
}

func (ba *BrasilApiService) FetchZipCode(zipCode string, startTime time.Time) (zipCodeInfo *entity.ZipCode, err error) {
	req, err := http.NewRequestWithContext(ba.Ctx, http.MethodGet, "https://brasilapi.com.br/api/cep/v1/"+zipCode, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ba.Ctx.Err() == context.DeadlineExceeded {
			return nil, errors.New("request time limit exceeded (brasilapi)")
		}
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
