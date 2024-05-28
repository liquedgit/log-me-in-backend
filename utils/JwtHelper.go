package utils

import (
	"log-me-in/jwt"
)

func GenerateNewToken(id string) (*string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["id"] = id
	secretkey, err := GetFromEnv("JWT_SECRET_KEY")
	if err != nil {
		return nil, err
	}
	tokenString, err := token.SignedString([]byte(*secretkey))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func DecodeToken(strToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) ([]byte, error) {
		secretkey, err := GetFromEnv("JWT_SECRET_KEY")
		if err != nil {
			return nil, err
		}
		return []byte(*secretkey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
