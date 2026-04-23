package main

import (
	"backend/handlers"
	"backend/repositories"
	"backend/services"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func initDB() *sql.DB {

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		dbHost = "db"
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := "disable"

	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUser, dbPass, dbName, dbSslMode)

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Berhasil konek ke DB")
				break
			}
		}
		fmt.Printf("Database belum siap, mencoba lagi dalam 2 detik... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Gagal konek ke DB setelah beberapa kali percobaan", err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR(20) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		link TEXT,
		is_email_verified BOOL NOT NULL,
		is_phone_verified BOOL NOT NULL,
		role VARCHAR(20) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		last_login TIMESTAMP
	)`

	_, err = db.Exec(createTableQuery)

	if err != nil {
		log.Fatal("Gagal Membuat Table", err)
	}

	fmt.Println("Database Siap")

	return db
}

func main() {

	db := initDB()
	defer db.Close()
	// depedency injection
	userRepo := repositories.NewUserPostgresRepositories(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/register", userHandler.RegisterHandler)

	router.Run(":8080")
}
