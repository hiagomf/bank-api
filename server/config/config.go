package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var config *Config

// GetConfig retorna um ponteiro para uma estrutura de configuração
// que contém todos os dados da configuração
func GetConfig() *Config {
	if config == nil {
		log.Fatal("a configuração não pode ser carregada")
	}
	return config
}

// LoadConfig carrega as configurações através do arquivo definido
func LoadConfig() {
	path := "/etc/bank-api/config.json"
	if val, set := os.LookupEnv("BANK_API_CONFIG"); set && val != "" {
		path = val
	} else {
		log.Println("variável de ambiente `BANK_API_CONFIG` não está definida, usado diretorio: ", path)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	if err := validateConfig(); err != nil {
		log.Fatal(err)
	}

	if config.SystemID == 0 {
		log.Fatal(errors.New("ID de sistema não configurado"))
	}
}

func validateConfig() error {
	if config == nil {
		return errors.New("a configuração não pode ser carregada")
	}

	if err := validateMainDatabase(); err != nil {
		return err
	}

	if config.ExternalAddress == "" {
		config.ExternalAddress = ":8081"
	}

	if config.ExternalPublicAddress == "" {
		config.ExternalPublicAddress = ":8082"
	}

	if config.InternalAddress == "" {
		config.InternalAddress = ":8083"
	}

	if config.AccessLogDirectory == "" {
		config.AccessLogDirectory = "/var/log/mub_brasil/access.log"
	}

	if config.ErrorLogDirectory == "" {
		config.ErrorLogDirectory = "/var/log/mub_brasil/error.log"
	}

	if config.QueryTimeout == 0.0 {
		config.QueryTimeout = 2.0
	}
	return nil
}

func validateMainDatabase() error {
	if len(config.Databases) == 0 {
		return errors.New("configuração de banco de dados não definida")
	}

	main := false
	for i := 0; i < len(config.Databases); i++ {
		if config.Databases[i].Main {
			if main {
				return errors.New("mais de um banco de dados foi definido como principal")
			}
			main = true
		}

		if config.Databases[i].TransactionTimeout == 0 {
			config.Databases[i].TransactionTimeout = 1
		}
	}

	if !main {
		return errors.New("é necessário definir um banco de dados principal")
	}

	return nil
}
