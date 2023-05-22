package main

import (
	"ass_3/pkg/store/postgres"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("DB_PORT")
	user := "web"
	password := "admin"
	dbname := "library"
	db, err := postgres.OpenDB(port, user, password, dbname)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

}
