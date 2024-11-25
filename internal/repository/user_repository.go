package repository

import (
    "database/sql"
    "cs-exp-go-api/internal/models"
    "fmt"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    rows, err := r.DB.Query("SELECT id, name, email, username FROM pengguna")
    if err != nil {
        return nil, fmt.Errorf("failed to query pengguna: %w", err)
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username); err != nil {
            return nil, fmt.Errorf("failed to scan user: %w", err)
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }

    return users, nil
}

func (r *UserRepository) Create(user models.User) error {
    _, err := r.DB.Exec("INSERT INTO pengguna (name, email, username, password) VALUES ($1, $2, $3, $4)",
        user.Name, user.Email, user.Username, user.Password)
    if err != nil {
        return fmt.Errorf("failed to insert pengguna: %w", err)
    }
    return nil
}

func (r *UserRepository) Update(user models.User) error {
    _, err := r.DB.Exec("UPDATE pengguna SET name = $1, email = $2, username = $3, password = $4 WHERE id = $5",
        user.Name, user.Email, user.Username, user.Password, user.ID)
    if err != nil {
        return fmt.Errorf("failed to update pengguna: %w", err)
    }
    return nil
}

func (r *UserRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM pengguna WHERE id = $1", id)
    if err != nil {
        return fmt.Errorf("failed to delete pengguna: %w", err)
    }
    return nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
    var user models.User
    row := r.DB.QueryRow("SELECT id, name, email, username, password FROM pengguna WHERE username = $1", username)
    if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to query pengguna by username: %w", err)
    }
    return &user, nil
}