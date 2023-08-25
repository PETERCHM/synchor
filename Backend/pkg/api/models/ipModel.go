// pkg/api/models/ipModel.go
package models

type IPInfo struct {
	ID        int    `json:"id"`
	IPAddress string `json:"ip_address"`
	Location  string `json:"location"`
}
