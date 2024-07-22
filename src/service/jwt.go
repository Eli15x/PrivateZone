package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	JWTSecretKey string
	// Adicione outras dependências conforme necessário
}

func NewAuthService(jwtSecretKey string) *AuthService {
	return &AuthService{
		JWTSecretKey: jwtSecretKey,
	}
}

func (s *AuthService) GenerateToken(userId string) (string, error) {
	// Gerar token JWT usando a chave secreta configurada
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 6).Unix(), // Token expira em 2 horas
	})

	// Assinar o token com a chave secreta e obter o token JWT como uma string
	signedToken, err := token.SignedString([]byte(s.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Outras funções do serviço de autenticação conforme necessário
