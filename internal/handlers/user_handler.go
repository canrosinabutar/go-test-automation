package handlers

import (
	"cs-exp-go-api/internal/models"
	"cs-exp-go-api/internal/services"
	"cs-exp-go-api/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
    Service *services.UserService
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    Pengguna, err := h.Service.GetAllUsers()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, Pengguna)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.Service.CreateUser(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.Service.UpdateUser(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.Service.DeleteUser(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.Service.CreateUser(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}


func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    // Decode the JSON request body
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Call the service to authenticate the user and generate a token
    user, token, err := h.Service.Login(credentials.Username, credentials.Password)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Create the response struct
    response := struct {
        User  *models.User `json:"user"`
        Token string       `json:"token"`
    }{
        User:  user,
        Token: token,
    }

    // Set the response content type to JSON
    w.Header().Set("Content-Type", "application/json")

    // Encode and send the response
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}