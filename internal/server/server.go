package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	mainRouting(r)

	err := http.ListenAndServe(":8080", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
