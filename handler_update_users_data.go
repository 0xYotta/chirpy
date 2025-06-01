package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/0xYotta/chirpy/internal/auth"
	"github.com/0xYotta/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpdateUsersData(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID          uuid.UUID `json:"id"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedAt   time.Time `json:"created_at"`
		Email       string    `json:"email"`
		IsChirpyRed bool      `json:"is_chirpy_red"`
	}

	type parameters struct {
		NewPassword string `json:"password"`
		NewEmail    string `json:"email"`
	}
	// email and password in body
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "token malformed or missing", err)
		return
	}
	userUUID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, 401, "token validation failed", err)
		return
	}

	defer r.Body.Close()

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	hashedPassword, err := auth.HashPassword(params.NewPassword)
	if err != nil {
		respondWithError(w, 500, "Can't hash the password", err)
	}

	if err := cfg.db.UpdatePasswordAndEmail(r.Context(), database.UpdatePasswordAndEmailParams{
		ID:             userUUID,
		Email:          params.NewEmail,
		HashedPassword: hashedPassword,
	}); err != nil {
		respondWithError(w, 500, "Couldn't update users data", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.NewEmail)
	if err != nil {
		respondWithError(w, 500, "Couldn't get user", err)
		return
	}

	respondWithJSON(w, 200, response{
		ID:          user.ID,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed.Bool,
	})
}
