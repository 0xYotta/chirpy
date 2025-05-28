package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}

	if !validateContentType(w, r) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
}

func validateContentType(w http.ResponseWriter, r *http.Request) bool {
	ct := r.Header.Get("Content-Type")

	if exist := strings.Contains(strings.ToLower(ct), "application/json"); !exist {
		log.Printf("Wrong Content-Type Header. got: %v", ct)
		respondWithError(w, 400, "Wrong Content-Type value", nil)
		return false
	}
	return true
}
