package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	vault "github.com/mittwald/vaultgo"
	"github.com/nasermirzaei89/env"
)

type VaultCred struct {
	Data struct {
		User     string `json:"username"`
		Password string `json:"password"`
	} `json:"data"`
}

type DbConnection struct {
	Dbname   string
	Host     string
	Port     int
	User     string
	Password string
}

func init() {
	_ = godotenv.Load("config/.env")
}

// Read (generate) credentials from our Vault server.
func getDBConnectionConfig() DbConnection {
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

// This function opens up a new Postgres connection to our server and returns it.
func openConnection() *sql.DB {
	config := getDBConnectionConfig()
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)

	//Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	//Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	err = db.PingContext(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to PostgreSQL db using user <%s> and password <%s>\n", config.User, config.Password)
	return db
}

func readAuthors(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			//Executing a call to recover inside a deferred function.
			if rcv := recover(); rcv != nil {
				log.Print("do not panic")
			}
		}()

		//Query executes a query that returns rows, typically a SELECT.
		rows, err := db.QueryContext(r.Context(), `SELECT author_id,first_name,last_name FROM authors`)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		names := []string{}
		//Next prepares the next result row for reading with the Scan method.
		for rows.Next() {
			var author_id, first_name, last_name string
			err = rows.Scan(&author_id, &first_name, &last_name)
			if err != nil {
				panic(err)
			}

			names = append(names, author_id, first_name, last_name)
		}

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprint(w, strings.Join(names, ", "))
		if err != nil {
			return
		}
	}
}

func main() {
	// Setup DB connection
	db := openConnection()
	defer db.Close()

	r := initRouting(db)

	hostPort := env.MustGetString("APP_HOST") + env.MustGetString("APP_PORT")
	fmt.Println("Listening on ", hostPort)
	err := http.ListenAndServe(hostPort, r)
	if err != nil {
		return
	}
}

func initRouting(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/read").HandlerFunc(readAuthors(db))
	return r
}
