package requests

import "fmt"

type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

func (r *BrasilAPIResponse) Address() string {
	return fmt.Sprintf(
		"Rua: %s; Bairro: %s; Cidade: %s; Estado: %s; CEP: %s",
		r.Street, r.Neighborhood, r.City, r.State, r.CEP)
}

type ViaCEPResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}

func (r *ViaCEPResponse) Address() string {
	return fmt.Sprintf(
		"Rua: %s; Bairro: %s; Cidade: %s; Estado: %s; CEP: %s",
		r.Street, r.Neighborhood, r.City, r.State, r.CEP)
}
