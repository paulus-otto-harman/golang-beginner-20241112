package main

import (
	"20241112/database"
	"20241112/router"
	"database/sql"
	"log"
	"net/http"
)

func DbProvider() *sql.DB {
	return database.DbOpen("20241111a")
}

func main() {
	db := database.DbOpen("20241112")
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
