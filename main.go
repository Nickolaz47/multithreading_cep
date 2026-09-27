package main

import (
	"context"
	"fmt"
	"multithreading_cep/internal/requests"
	"time"
)

type AddressResponse interface {
	Address() string
}

type message struct {
	API      string
	Response AddressResponse
}

func worker(ctx context.Context, URL, apiName string,
	out AddressResponse, ch chan message) {
	err := requests.DoGetRequest(ctx, URL, out)
	if err == nil {
		ch <- message{API: apiName, Response: out}
	}
	close(ch)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cep := "89283132"
	brasilAPIURL := requests.BuildBrasilApiURL(cep)
	viaCepURL := requests.BuildViaCepURL(cep)

	brasilCh := make(chan message)
	viaCh := make(chan message)

	go worker(ctx, brasilAPIURL, "BrasilAPI", &requests.BrasilAPIResponse{},
		brasilCh)
	go worker(ctx, viaCepURL, "ViaCEP", &requests.ViaCEPResponse{}, viaCh)

	select {
	case msg := <-brasilCh:
		fmt.Printf("Mensagem recebida de %s com o conteúdo %v\n", msg.API,
			msg.Response.Address())
	case msg := <-viaCh:
		fmt.Printf("Mensagem recebida de %s com o conteúdo %v\n", msg.API,
			msg.Response.Address())
	case <-ctx.Done():
		fmt.Println("Ambas as APIs demoraram demais.")
	}
}
