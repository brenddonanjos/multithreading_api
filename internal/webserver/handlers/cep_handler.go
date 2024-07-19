package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brenddonanjos/multithreading_api/internal/app/service"
	"github.com/go-chi/chi/v5"
)

type CepHandler struct{}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

func (h *CepHandler) GetCepInfo(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		fmt.Println("Cep is required")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Cep is required")
		return
	}

	cepInfo, err := service.GetZipCodeInfo(cep)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	fmt.Println("==============")
	fmt.Println(
		"Cep: "+cepInfo.ZipCode,
		"\nState: "+cepInfo.State,
		"\nCity: "+cepInfo.City,
		"\nNeighborhood: "+cepInfo.Neighborhood,
		"\nStreet: "+cepInfo.Street,
		"\nService: "+cepInfo.Service,
		"\nExecution Time: "+cepInfo.ExecutionTime,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cepInfo)

}
