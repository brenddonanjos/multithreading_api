package entity

type ZipCode struct {
	ZipCode       string `json:"zip_code"`
	State         string `json:"state"`
	City          string `json:"city"`
	Neighborhood  string `json:"neighborhood"`
	Street        string `json:"street"`
	Service       string `json:"service"`
	ExecutionTime string `json:"execution_time"`
}

func NewZipCode(cep, state, city, neighborhood, street, service string, executionTime string) *ZipCode {
	return &ZipCode{
		ZipCode:       cep,
		State:         state,
		City:          city,
		Neighborhood:  neighborhood,
		Street:        street,
		Service:       service,
		ExecutionTime: executionTime,
	}
}
