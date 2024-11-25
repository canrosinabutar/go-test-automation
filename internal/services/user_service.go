package services

import (
    "cs-exp-go-api/internal/models"
    "cs-exp-go-api/internal/repository"
    "errors"
    "log"
    "time"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    Repo *repository.UserRepository
    JwtSecret string
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetAll()
}

func (s *UserService) CreateUser(user models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    log.Printf("Creating user: %s", user.Username)
    log.Printf("Hashed Password: %s", string(hashedPassword))
    user.Password = string(hashedPassword)
    return s.Repo.Create(user)
}

func (s *UserService) UpdateUser(user models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    log.Printf("Updating user: %s", user.Username)
    log.Printf("Hashed Password: %s", string(hashedPassword))
    user.Password = string(hashedPassword)
    return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
    log.Printf("Deleting user with ID: %d", id)
    return s.Repo.Delete(id)
}

func (s *UserService) Login(username, password string) (*models.User, string, error) {
    // Retrieve the user from the repository
    user, err := s.Repo.GetByUsername(username)
    if err != nil {
        return nil, "", err
    }
    if user == nil {
        // User not found
        return nil, "", errors.New("user not found")
    }

    // Compare the provided password with the stored hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        // Password mismatch
        return nil, "", errors.New("invalid password")
    }

    // Create JWT token
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
        Username: user.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(s.JwtSecret))
    if err != nil {
        return nil, "", err
    }
    

    return user, tokenString, nil
}