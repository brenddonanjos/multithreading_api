package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/brenddonanjos/multithreading_api/internal/entity"
)

type ViaCepService struct{}

func NewViaCepService() *ViaCepService {
	return &ViaCepService{}
}

func (vc *ViaCepService) FetchZipCode(zipCode string, startTime time.Time) (zipCodeInfo *entity.ZipCode, err error) {
	req, err := http.NewRequest(http.MethodGet, "https://viacep.com.br/ws/"+zipCode+"/json/", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
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
