package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/PETERCHM/synchor/backend/pkg/api/models"
	"github.com/PETERCHM/synchor/backend/pkg/repositories"
	"github.com/PETERCHM/synchor/backend/pkg/services"
)

func GetIPInfoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup

		dataChannel := make(chan models.IPInfo)

		wg.Add(1)
		go repositories.FetchIPData(db, dataChannel, &wg)

		data, ok := <-dataChannel
		if !ok {
			http.Error(w, "Data channel closed", http.StatusInternalServerError)
			return
		}

		wg.Add(1)
		go services.ProcessIPData(data, &wg)

		wg.Wait()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Data synchronization complete"}`))
	}
}
