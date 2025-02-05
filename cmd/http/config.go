package main

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

// GetLevelLog converte o valor do nível de log em um tipo `slog.Level`.
func (c Config) GetLevelLog() slog.Level {
	switch strings.ToLower(c.LevelLog) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Config é uma estrutura que representa as configurações do ambiente.
// Cada campo possui tags para indicar o nome da variável de ambiente associada
// e informações adicionais, como valor padrão ou se é obrigatório.
type Config struct {
	TestLoadEnv string `env:"TEST_LOAD_ENV,info"`    // Variável de ambiente opcional para teste
	ServerPort  string `env:"SERVER_PORT, 5000"`     // Porta padrão do servidor (5000 se não especificada)
	DbConnUrl   string `env:"DB_CONN_URL, required"` // URL database
	LevelLog    string `env:"LEVEL_LOG,info"`        // Nível de log (padrão: info)
}

// SPrint retorna para o client uma string com as variáveis de ambiente e seus valores padrão.
func (c Config) SPrint() (envs string) {
	v := reflect.ValueOf(c) // Obtém o valor da estrutura
	t := v.Type()           // Obtém o tipo da estrutura
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)                                // Obtém o campo atual
		envTag := strings.Split(field.Tag.Get("env"), ",") // Obtém a tag "env"
		name := envTag[0]                                  // Nome da variável de ambiente
		value := envTag[1]                                 // Valor padrão ou flag
		envs += fmt.Sprintf("%s - %s\n", name, value)
	}
	return
}

// loadFromEnv carrega os valores das variáveis de ambiente para a estrutura Config.
func (c *Config) loadFromEnv() {
	v := reflect.ValueOf(c).Elem() // Referência para o valor da estrutura
	t := v.Type()                  // Tipo da estrutura

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)                                // Campo atual da estrutura
		envTag := strings.Split(field.Tag.Get("env"), ",") // Tag "env"
		envName := strings.TrimSpace(envTag[0])            // Nome da variável de ambiente
		value := strings.TrimSpace(os.Getenv(envName))     // Valor da variável de ambiente

		// Verifica se há um valor padrão na tag
		var defaultValue string
		if len(envTag) > 1 {
			defaultValue = strings.TrimSpace(envTag[1])
		}

		// Usa o valor da variável de ambiente ou o valor padrão, se disponível
		if value == "" {
			if defaultValue != "required" { // Não preenche com valor padrão se "required"
				value = defaultValue
			}
		}

		// Atualiza o campo na struct
		v.FieldByName(field.Name).SetString(value)
	}
}

// validate verifica se todas as variáveis obrigatórias possuem valores válidos.
func (c Config) validate() {
	var validationMsg string
	v := reflect.ValueOf(c) // Referência ao valor da estrutura
	t := v.Type()           // Tipo da estrutura

	for i := 0; i < v.NumField(); i++ {
		value := strings.TrimSpace(v.Field(i).String())         // Valor do campo
		envTag := strings.Split(t.Field(i).Tag.Get("env"), ",") // Tag "env"
		envName := envTag[0]                                    // Nome da variável de ambiente
		envValue := strings.TrimSpace(envTag[1])                // Valor padrão ou flag

		if envValue == "required" && value == "" { // Verifica se o campo obrigatório está vazio
			validationMsg += fmt.Sprintf("%s is required\n", envName)
		}
	}

	// Lança um erro se houver mensagens de validação
	if len(validationMsg) != 0 {
		panic(validationMsg)
	}
}

// loadConfig carrega as configurações a partir de variáveis de ambiente e valida.
func loadConfig() Config {
	// Carrega as variáveis de ambiente de um arquivo .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := Config{}
	config.loadFromEnv() // Carrega variáveis de ambiente.
	config.validate()    // Valida os valores.

	return config
}
