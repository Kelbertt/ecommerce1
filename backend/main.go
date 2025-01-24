package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Backend funcionando!")
    })

    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}