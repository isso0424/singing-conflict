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
	http.ListenAndServe(":80", nil)
}
