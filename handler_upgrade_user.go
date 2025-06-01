package main

import (
	"encoding/json"
	"net/http"

	"github.com/0xYotta/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeUserToRed(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, "no auth header", err)
		return
	}
	if apiKey != cfg.polkaAPIKey {
		respondWithError(w, 401, "invalid api key", err)
		return
	}

	defer r.Body.Close()
	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, 500, "Couldn't decode r.Body", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, 204, "", nil)
		return
	}

	userUUID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, 500, "couldn't parse uuid", err)
		return
	}

	if err := cfg.db.UpgradeUserToRed(r.Context(), userUUID); err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
