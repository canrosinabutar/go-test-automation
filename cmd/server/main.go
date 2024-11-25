package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "cs-exp-go-api/internal/router"
)

func main() {
    // Retrieve database connection details from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    // Construct the connection string
    connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

    // Connect to the database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    defer db.Close()

    // Retrieve JWT secret from environment variable
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET environment variable not set")
    }

    // Create router with db and JWT secret
    r := router.NewRouter(db, jwtSecret)

    // Start server
    log.Fatal(http.ListenAndServe(":8989", r))
}