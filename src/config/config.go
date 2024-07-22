// config.go

package config

import (
	"os"
)

// Config estrutura para armazenar configurações da aplicação
type Config struct {
	JWTSecretKey string
	// Adicione outras configurações conforme necessário
}

var config *Config

// GetConfig retorna a instância única da configuração da aplicação
func GetConfig() *Config {
	return config
}

// LoadConfig carrega as configurações da aplicação
func LoadConfig() {
	// Aqui você pode carregar as configurações da maneira que preferir
	// Exemplo simples: lendo variáveis de ambiente ou de um arquivo de configuração

	// Carregar chave secreta JWT (exemplo com variável de ambiente)
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	// Preencher a estrutura de configuração
	config = &Config{
		JWTSecretKey: jwtSecretKey,
		// Adicione outras configurações aqui
	}
}
