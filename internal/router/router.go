package router

import (
    "database/sql"
    "cs-exp-go-api/internal/handlers"
    "cs-exp-go-api/internal/repository"
    "cs-exp-go-api/internal/services"
    "cs-exp-go-api/internal/middleware"
    "net/http"

    "github.com/gorilla/mux"
)

func NewRouter(db *sql.DB, jwtSecret string) *mux.Router {
    r := mux.NewRouter()

    userRepo := &repository.UserRepository{DB: db}
    userService := &services.UserService{Repo: userRepo, JwtSecret: jwtSecret} // Make sure JwtSecret is set in the service
    userHandler := &handlers.UserHandler{Service: userService}

    // Public routes
    r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
    r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

    // Protected routes with JWT middleware
    r.Handle("/users", middleware.JWTAuthMiddleware(jwtSecret)(http.HandlerFunc(userHandler.GetAllUsers))).Methods("GET")
    r.Handle("/users", middleware.JWTAuthMiddleware(jwtSecret)(http.HandlerFunc(userHandler.CreateUser))).Methods("POST")
    r.Handle("/users/{id:[0-9]+}", middleware.JWTAuthMiddleware(jwtSecret)(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")

    // Update route should be specific and not overlap with create or list routes
    r.Handle("/users/{id:[0-9]+}", middleware.JWTAuthMiddleware(jwtSecret)(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")

    return r
}