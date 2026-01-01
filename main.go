package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {
	fmt.Println("We are running!")

	var err error
	DB, err = sql.Open("sqlite3", "./pokedexDB.db")
	if err != nil {
		log.Fatal(nil)
	}
	defer DB.Close()

	//mux := http.NewServeMux()
	r := mux.NewRouter()

	r.Handle("/", &homeHandler{})
	r.HandleFunc("/pokemon", allPokemonHandler).Methods("GET")
	r.HandleFunc("/pokemon/{id}", getPokemonHandler).Methods("GET")
	r.HandleFunc("/pokemon", createPokemonHandler).Methods("POST")
	r.HandleFunc("/pokemon/{id}", updatePokemonHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
