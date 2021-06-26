package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func getRandomNumber(min int, max int) int {
	return (rand.Intn(max-min) + min)
}

func main() {
	config := readConfig()

	listenAt := fmt.Sprintf(":%d", config.port)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		responseFormat := req.URL.Query().Get("format")

		names := strings.Split(req.URL.Query().Get("names"), ",")

		randomIndex := getRandomNumber(0, len(names))

		randomName := names[randomIndex]
		fmt.Println(randomName)

		if responseFormat == "json" {
			res.Header().Set("Content-type", "application/json;charset=utf-8")

			json.NewEncoder(res).Encode(struct {
				Name string `json:"name"`
			}{
				Name: randomName,
			})

			return
		}

		fmt.Fprint(res, randomName)

	})

	log.Printf("Open the following URL in the browser: http://localhost:%d\n", config.port)

	log.Fatal(http.ListenAndServe(listenAt, nil))
}
