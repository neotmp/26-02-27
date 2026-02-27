package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Postgres: unable to load .env file: %e", err)
	}

	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))

	// Connect to database
	database, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	//defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	fmt.Println("Connected to Postgres")

	DB = database

}
