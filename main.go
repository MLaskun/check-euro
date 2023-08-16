package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const Host = "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json"

func main() {
	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 10; i++ {
		now := time.Now().Format(time.RFC1123)
		start := time.Now()

		req, err := http.NewRequest(http.MethodGet, Host, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("User-Agent", "Inflacja rosnie")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		responseTime := fmt.Sprintf("[%s] Time from request to response: %s", now, time.Since(start))
		log.Println(responseTime)
		statusCode := fmt.Sprintf("[%s] Status code: %d", now, resp.StatusCode)
		log.Println(statusCode)
		var isJson string
		var jsonValid string

		for _, value := range resp.Header["Content-Type"] {
			if value == "application/json; charset=utf-8" {
				isJson = fmt.Sprintf("[%s] Response content type is JSON", now)
				log.Println(isJson)
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				jsonValid = fmt.Sprintf("[%s] Is json valid: %t", now, json.Valid(body))
				log.Println(jsonValid)
			} else {
				isJson = fmt.Sprintf("[%s] Response content type is not JSON but: %s", now, value)
				log.Println(isJson)
				jsonValid = fmt.Sprintf("[%s] There is no JSON in response", now)
				log.Println(jsonValid)
			}
		}

		f.WriteString(responseTime + "\n" + statusCode + "\n" + isJson + "\n" + jsonValid + "\n")

		if i < 9 {
			time.Sleep(5 * time.Second)
		}
	}

}
