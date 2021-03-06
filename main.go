package main

import (
	"isso0424/singing-conflict/server/handler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/webhook", handler.Handler)
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
