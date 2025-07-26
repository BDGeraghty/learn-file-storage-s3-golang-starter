package main

import (
	"net/http"
	
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment."))
		return
	}

	err := cfg.db.Reset()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't reset database", err)
		return
	}
	
	// Clear the in-memory thumbnail storage
	videoThumbnails = make(map[uuid.UUID]thumbnail)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database reset to initial state"))
}
