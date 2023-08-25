// pkg/api/handlers/ipHandlers.go
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/astaxie/beego"	
	"github.com/prometheus/common/model"
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

		// Respond with success message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Data synchronization complete"}`))
	}
}
