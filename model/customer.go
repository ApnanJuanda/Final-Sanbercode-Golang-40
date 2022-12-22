package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var signature = []byte("myPrivateSignateure")

type Customer struct {
	ID               int       `json:"id,omitempty"`
	Email            string    `json:"email,omitempty"`
	Password         string    `json:"password,omitempty"`
	Name             string    `json:"name,omitempty"`
	PhoneNumber      string    `json:"phone_number,omitempty"`
	Address          string    `json:"address,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `json:"verified"`
}

type CustomerRegister struct {
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
}

type CustomerLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *Customer) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"customer_id": u.ID,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
		"iss":         "sanbercode",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signature)
	return stringToken, err
}

func (u *Customer) CorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *Customer) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return signature, nil
	})
	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}

	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	fmt.Println("Decript JWT", parsedToken.Claims.(jwt.MapClaims))
	return parsedToken.Claims.(jwt.MapClaims), nil
}
