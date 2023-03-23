package main

import (
	"encoding/json"

	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type Produto struct {
	ID    string    
	Nome  string 
	Preco float64
}

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.Use(handlers.CORS(headers, methods, origins))

	router.HandleFunc("/adicionar-item/{idProduto}", func(w http.ResponseWriter, requisicao *http.Request) {
		vars := mux.Vars(requisicao)
		idProduto := vars["idProduto"]

		produtos := []Produto{
			Produto{ID: "produto1", Nome: "Produto 1", Preco: 19.99},
			Produto{ID: "produto2", Nome: "Produto 2", Preco: 29.99},
			Produto{ID: "produto3", Nome: "Produto 3", Preco: 39.99},
			Produto{ID: "produto4", Nome: "Produto 4", Preco: 49.99},
			Produto{ID: "produto5", Nome: "Produto 5", Preco: 59.99},
		}

		var produtoDesejado Produto
		for _, produto := range produtos {
			if produto.ID == idProduto {
				produtoDesejado = produto
				break
			}
		}

		if produtoDesejado.ID != "" {
			jsonData, err := json.Marshal(produtoDesejado)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else {
			http.Error(w, "Produto não encontrado", http.StatusNotFound)
		}
	}).Methods("POST")

	log.Println("Servidor iniciado na porta 3333")
	log.Fatal(http.ListenAndServe(":3333", router))
}
