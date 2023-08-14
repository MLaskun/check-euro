package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const Host = "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json"

func main() {
	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, Host, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "Inflacja rosnie")

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("Time from request to response:", time.Since(start))
	log.Println("Status code:", resp.StatusCode)

	for _, value := range resp.Header["Content-Type"] {
		if value == "application/json; charset=utf-8" {
			log.Println("Response content type is JSON")
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Is json valid:", json.Valid(body))
		} else {
			log.Printf("Response content type is not JSON but: %s", value)
			log.Println("There is no JSON in response")
		}
	}

}
