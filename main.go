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

	now := time.Now().Format(time.RFC1123)

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
	log.Printf("[%s] Time from request to response: %s", now, time.Since(start))
	log.Printf("[%s] Status code: %d", now, resp.StatusCode)

	for _, value := range resp.Header["Content-Type"] {
		if value == "application/json; charset=utf-8" {
			log.Printf("[%s] Response content type is JSON", now)
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("[%s] Is json valid: %t", now, json.Valid(body))
		} else {
			log.Printf("[%s] Response content type is not JSON but: %s", now, value)
			log.Printf("[%s] There is no JSON in response", now)
		}
	}

}
