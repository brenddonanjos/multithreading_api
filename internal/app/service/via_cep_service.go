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

type ViaCepService struct {
	Ctx context.Context
}

func NewViaCepService(ctx context.Context) *ViaCepService {
	return &ViaCepService{
		Ctx: ctx,
	}
}

func (vc *ViaCepService) FetchZipCode(zipCode string, startTime time.Time) (zipCodeInfo *entity.ZipCode, err error) {
	req, err := http.NewRequestWithContext(vc.Ctx, http.MethodGet, "https://viacep.com.br/ws/"+zipCode+"/json/", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if vc.Ctx.Err() == context.DeadlineExceeded {
			return nil, errors.New("request time limit exceeded (viacep)")
		}
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	viaCep := &entity.ViaCep{}
	err = json.Unmarshal(body, &viaCep)
	if err != nil {
		return nil, err
	}
	if viaCep.Cep == "" {
		return nil, errors.New("cep not found (viacep)")
	}

	executionTime := time.Since(startTime).String()

	zipCodeInfo = entity.NewZipCode(viaCep.Cep, viaCep.Uf, viaCep.Localidade, viaCep.Bairro, viaCep.Logradouro, "Via CEP", executionTime)
	return zipCodeInfo, nil

}
