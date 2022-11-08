package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	cfgVault "vault-psql/internal/pkg/vault"
)

//GetPostgresqlConfig get information config postgresql from vault secret dynamic.
func GetPostgresqlConfig() *sql.DB {
	config := cfgVault.GetVaultConfig()
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname,
	)

	//Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	//db.SetConnMaxIdleTime()
	//db.SetConnMaxLifetime()
	//db.SetMaxIdleConns()
	//db.SetMaxOpenConns()

	//Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	err = db.PingContext(context.Background())
	if err != nil {
		panic(err)
	}

	return db
}
