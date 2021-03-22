package main

import (
	"isso0424/singing-conflict/server/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handler.Handler)
	http.ListenAndServe(":4000", nil)
}
