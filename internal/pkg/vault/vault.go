package vault

import (
	vault "github.com/mittwald/vaultgo"
	"github.com/nasermirzaei89/env"
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
	client, errVault := vault.NewClient(env.MustGetString("VAULT_HOST")+env.MustGetString("VAULT_PORT"), vault.WithCaPath(""), vault.WithAuthToken(env.MustGetString("VAULT_TOKEN")))
	if errVault != nil {
		panic(errVault)
	}

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
