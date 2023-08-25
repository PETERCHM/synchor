// pkg/services/ipService.go
package services

import (
	"fmt"
	"sync"

	"github.com/PETERCHM/synchor/backend/pkg/api/models"
)

func ProcessIPData(data models.IPInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate data processing
	result := fmt.Sprintf("%s - Processed", data.Location)
	fmt.Println(result)
}
