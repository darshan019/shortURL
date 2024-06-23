package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func godot_env_var(key string) string {
	ok := godotenv.Load(".env")
	if ok != nil {
		fmt.Println("error loading env")
	}
	return os.Getenv(key)
}

func Initialize() {
	// connstr := "postgres://postgres:darshan123@localhost:5432/ShortURL?sslmode=disable"

	connstr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", godot_env_var("DBUSERNAME"), godot_env_var("DBPASSWORD"), godot_env_var("PORT"), godot_env_var("DBNAME"))

	var err error
	db, err = sql.Open("postgres", connstr)

	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging to db: %v", err)
	}

	fmt.Println("Connected to db")

	create_tables()
}

func create_tables() {
	create_user_table := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`

	create_url_table := `
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			short_url TEXT NOT NULL,
			long_url TEXT NOT NULL
		);
	`

	if _, err := db.Exec(create_user_table); err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}

	if _, err := db.Exec(create_url_table); err != nil {
		log.Fatalf("Error creating URL table: %v", err)
	}

	fmt.Println("Tables created successfully")
}

func Get_DB() *sql.DB { return db }
