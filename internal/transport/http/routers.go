package dockertest_test

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vault-psql/internal/service/author"
)

type ConfigService struct {
	Port      string
	AuthorSvc author.Service
}

func HandlerRequest(cfg ConfigService) error {
	authorHndlr := newAuthorHandler(cfg.AuthorSvc)

	// NewRouter returns a new router instance.
	r := mux.NewRouter()

	r.HandleFunc("/authors", authorHndlr.HandlerListAuthors).Methods(http.MethodGet)
	r.HandleFunc("/author", authorHndlr.HandlerCreateAuthor).Methods(http.MethodPost)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("start application on listening on %s", cfg.Port)

		if err := http.ListenAndServe(cfg.Port, r); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	time.Sleep(1 * time.Second)
	log.Printf("shutting down application ...\n")

	return nil
}
