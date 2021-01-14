package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func HSHA256Handler(password string) string {
	s := os.Getenv("SECRET_KEY")

	h := hmac.New(sha256.New, []byte(s))
	h.Write([]byte(password))

	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return sha
}

func TokenGenerator(subject string, minute int) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(minute)).Unix(),
		Subject:   subject,
	}

	SIGN_KEY := []byte(os.Getenv("SIGN_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SIGN_KEY)
}
