package requests

func BuildBrasilApiURL(cep string) string {
	return "https://brasilapi.com.br/api/cep/v1/" + cep
}

func BuildViaCepURL(cep string) string {
	return "https://viacep.com.br/ws/" + cep + "/json"
}
