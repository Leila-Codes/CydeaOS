package media

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrInvalidMediaID = errors.New("invalid media ID")
)

// Service - HTTP REST-ful(ish) endpoint for streaming of media files
func Service(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID, parseError := uuid.Parse(vars["mediaID"])

	if parseError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrInvalidMediaID.Error()))
		return
	}

	if entry, hasEntry := definitions[mediaID]; hasEntry {
		http.ServeFile(w, r, entry.Path)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(ErrInvalidMediaID.Error()))
	}
}
