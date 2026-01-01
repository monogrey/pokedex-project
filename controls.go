package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// gets all pokemon in list(database)
func allPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var pokemons = make([]Pokemon, 0)

	rows, err := DB.Query("SELECT * FROM pokemon")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	for rows.Next() {
		var p Pokemon

		err = rows.Scan(&p.ID, &p.Name, &p.Desc)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}

		pokemons = append(pokemons, p)
	}
	json.NewEncoder(w).Encode(pokemons)
}

// GET function for web application
func getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	pokemonID, err := strconv.Atoi(idStr)

	pokemon, err := getPokemon(DB, pokemonID)
	if err != nil {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func getPokemon(db *sql.DB, id int) (*Pokemon, error) {
	query := "SELECT * FROM pokemon WHERE id=?"
	rows := db.QueryRow(query, id)

	pokemon := &Pokemon{}
	err := rows.Scan(&pokemon.ID, &pokemon.Name, &pokemon.Desc)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

// POST function for web application
func createPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var p Pokemon
	json.NewDecoder(r.Body).Decode(&p)

	err := createPokemon(DB, p.Name, p.Desc)
	if err != nil {
		http.Error(w, "Failed to create pokemon", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Pokemon created successfully!")
}

func createPokemon(db *sql.DB, name, desc string) error {
	query := "INSERT INTO pokemon(desc) VALUES (?)"
	_, err := db.Exec(query, name, desc)

	if err != nil {
		return nil
	}
	return nil
}

// PUT function for web application
func updatePokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	pokemonID, err := strconv.Atoi(idStr)

	var p Pokemon
	err = json.NewDecoder(r.Body).Decode(&p)

	updatePokemon(DB, pokemonID, p.Name, p.Desc)
	if err != nil {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Pokemon updated successfully!")
}

func updatePokemon(db *sql.DB, id int, name, desc string) error {
	query := "UPDATE pokemon SET name = ?, desc = ? WHERE id = ?"
	_, err := db.Exec(query, name, desc, id)
	if err != nil {
		return err
	}
	return nil
}
