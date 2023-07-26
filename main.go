package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json", nil)
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
	log.Println("Response content type:", resp.Header["Content-Type"])

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Is json:", json.Valid(body))
}
