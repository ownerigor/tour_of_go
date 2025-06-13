package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Olá, você acessou o endpoint %s", req.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor rodando em http://localhost:8000")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
