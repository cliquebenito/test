package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"test/internal/models"
	"test/pkg/repository"
	"time"
)

const (
	salt       = "fajsfsjaj"
	tokenTTL   = 12 * time.Hour
	signingKey = "312fassafsasa$123"
)

type AuthService struct {
	repo repository.Auth
}

type tokenClaimts struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "User not found", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimts{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}
func (s *AuthService) ParseToken(token string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaimts{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := parsedToken.Claims.(*tokenClaimts)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
