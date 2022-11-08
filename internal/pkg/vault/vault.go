package vault

import (
	vault "github.com/mittwald/vaultgo"
	"github.com/nasermirzaei89/env"
	"github.com/pkg/errors"
	"strconv"
)

type DbConnection struct {
	Dbname   string
	Host     string
	Port     int
	User     string
	Password string
}

type VaultCred struct {
	Data struct {
		User     string `json:"username"`
		Password string `json:"password"`
	} `json:"data"`
}

//GetVaultConfig Read (generate) credentials from our Vault server.
func GetVaultConfig() DbConnection {
	client := initVault()

	key := []string{"v1", "database", "creds", env.MustGetString("VAULT_ROLE")}
	options := &vault.RequestOptions{}
	response := &VaultCred{}

	errRead := client.Read(key, response, options)
	if errRead != nil {
		panic(errRead)
	}

	port, errAtoi := strconv.Atoi(env.MustGetString("PGSQL_PORT"))
	if errAtoi != nil {
		panic(errAtoi)
	}

	return DbConnection{
		Dbname:   env.MustGetString("PGSQL_DB_NAME"),
		Host:     env.MustGetString("PGSQL_HOST"),
		Port:     port,
		User:     response.Data.User,
		Password: response.Data.Password,
	}
}

//initVault  returns a vault configuration for the client. It is safe to modify the return value of this function.
func initVault() *vault.Client {
	client, err := vault.NewClient(env.MustGetString("VAULT_HOST")+env.MustGetString("VAULT_PORT"), vault.WithCaPath(""), vault.WithAuthToken(env.MustGetString("VAULT_TOKEN")))
	if err != nil {
		panic(errors.Wrap(err, "error return config vault"))
	}
	return client
}
