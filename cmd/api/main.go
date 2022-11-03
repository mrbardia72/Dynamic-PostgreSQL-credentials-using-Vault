package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq" // import postgres driver
	"github.com/nasermirzaei89/env"
	cfg "vault-psql/internal/pkg/postgresql"
	repox "vault-psql/internal/repository/postgresql"
	"vault-psql/internal/service/author"
	router "vault-psql/internal/transport/http"
)

func main() {
	infoDB := cfg.GetPostgresqlConfig()
	defer func() { _ = infoDB.Close() }()

	authorRepo := repox.NewAuthorRepository(infoDB)
	authorSvc := author.NewService(authorRepo)

	rcs := router.ConfigService{
		Port:      env.MustGetString("APP_HOST") + env.MustGetString("APP_PORT"),
		AuthorSvc: authorSvc,
	}

	err := router.HandlerRequest(rcs)
	if err != nil {
		panic(fmt.Sprintf("error on handle http requests: %s", err))
	}
}
