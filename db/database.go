package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Create the schedule table if it doesn't exist
	_, err = DB.Exec(`
    CREATE TYPE status_enum AS ENUM ('ongoing', 'success', 'failed');
    
    CREATE TABLE IF NOT EXISTS schedule (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        start_time TIMESTAMP NOT NULL,
        end_time TIMESTAMP NOT NULL,
        status status_enum NOT NULL
    );
`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("Connected to PostgreSQL and ensured table exists.")
}
