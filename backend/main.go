package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Inicializa o roteador
	r := mux.NewRouter()

	// Definir uma rota simples
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mensagem do Backend, está funcionando."))
	})

	// Configuração CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Permitir apenas o frontend do Vite
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Aplica o middleware CORS
	handler := c.Handler(r)

	// Inicia o servidor
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
