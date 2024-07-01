package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/andrew-nino/em_time-tracker/entity"
	repo "github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// TODO  Перекинуть в config
const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repo.Authorization
}

func NewAuthService(repo repo.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Hashes the password and transfers the data to the repository.
func (s *AuthService) CreateUser(user entity.User) (int, error) {
	var err error
	user.PassportSerie, err = generatePasswordHash(user.PassportSerie)
	if err != nil {
		return 0, err
	}
	user.PassportNumber, err = generatePasswordHash(user.PassportNumber)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateUser(user)
}

// Checks that the client is already registered and returns the generated token.
func (s *AuthService) SignIn(serie, number string) (string, error) {
	serieHash, err := generatePasswordHash(serie)
	if err != nil {
		return "", err
	}
	numberHhash, err := generatePasswordHash(number)
	if err != nil {
		return "", err
	}
	user, err := s.repo.GetUser(serieHash, numberHhash)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// generatePasswordHash generates a SHA1 hash of the given password with a salt.
// The salt is a constant string.
func generatePasswordHash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		logrus.Debugf("failed to generate password hash: %s", err)
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}
