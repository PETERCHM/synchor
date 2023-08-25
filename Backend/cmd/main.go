// cmd/main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kevinburke/handlers"	
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/astaxie/beego"	
	"github.com/prometheus/common/model"
)

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/get-ip", handlers.GetIPInfoHandler(db))

	port := "8080"
	fmt.Printf("Server listening on port %s...\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
