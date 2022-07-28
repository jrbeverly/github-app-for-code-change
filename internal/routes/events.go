package routes

import (
	"net/http"

	"github.com/jrbeverly/github-app-for-code-change/internal/storage"
)

func eventHandler(w http.ResponseWriter, r *http.Request) {
	switch e := middleware.Payload.(type) {
	case storage.TestTriggerEvent:
		storage.PerformTestTrigger(e)
	case storage.ConfigChangeEvent:
		storage.PerformConfigTrigger(e)
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
