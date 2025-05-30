package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	db := cfg.db
	dbChirps, err := db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "error occured while fetching chirps from DB", err)
		return
	}

	var res []Chirp

	for _, chirp := range dbChirps {
		res = append(res, Chirp{
			chirp.ID,
			chirp.CreatedAt,
			chirp.UpdatedAt,
			chirp.UserID,
			chirp.Body,
		})
	}

	respondWithJSON(w, 200, res)
}

func (cfg *apiConfig) handlerGetChirpById(w http.ResponseWriter, r *http.Request) {
	db := cfg.db
	id := r.PathValue("chirpID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, 400, "error: not uuid", err)
	}
	chirp, err := db.GetChirpById(r.Context(), uuid)
	if err != nil {
		respondWithError(w, 404, "error: chirp not found", err)
	}

	respondWithJSON(w, 200, Chirp{
		chirp.ID,
		chirp.CreatedAt,
		chirp.UpdatedAt,
		chirp.UserID,
		chirp.Body,
	})
}
