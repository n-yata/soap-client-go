package main

import (
	"fmt"

	"log"

	"github.com/hooklift/gowsdl/soap"
)

func main() {
	// 1. SOAPクライアントの初期化
	client := soap.NewClient("http://localhost:8080/ws") // WSDLのエンドポイント

	// 2. CountriesPortサービスの作成
	service := NewCountriesPort(client)

	// 3. リクエストデータの作成
	request := &GetCountryRequest{
		Name: "Spain",
	}

	// 4. SOAP操作の呼び出し
	response, err := service.GetCountry(request)
	if err != nil {
		log.Fatalf("Error calling GetCountry: %v", err)
	}

	// 5. レスポンスの表示
	fmt.Printf("Country: %s\n", response.Country.Name)
	fmt.Printf("Population: %d\n", response.Country.Population)
	fmt.Printf("Capital: %s\n", response.Country.Capital)
	fmt.Printf("Currency: %s\n", *response.Country.Currency)
}
