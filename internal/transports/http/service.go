package http

import (
	"encoding/json"
	"github.com/fwidjaya20/demo-go-grpc-server/internal"
	"github.com/fwidjaya20/demo-go-grpc-server/internal/models"
	"net/http"
)

func Transfer(transferService internal.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Transfer

		err := json.NewDecoder(r.Body).Decode(&payload)

		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := transferService.Transfer(&payload)

		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(result)
	}
}

func GetHistory(transferService internal.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		history, err := transferService.GetHistory()

		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, _ := json.Marshal(history)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(res)
	}
}
