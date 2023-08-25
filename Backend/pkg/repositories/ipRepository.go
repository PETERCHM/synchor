// pkg/repositories/ipRepository.go
package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/PETERCHM/synchor/backend/pkg/api/models"
)

func FetchIPData(db *sql.DB, ch chan<- models.IPInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate fetching data from a database or API
	ip := "127.0.0.1" // Simulated IP input
	// ... implement SQL query or API call to fetch data ...
	data := models.IPInfo{
		ID:        1,
		IPAddress: ip,
		Location:  "Localhost",
	}

	ch <- data
	close(ch)
}
