package main

import (
	"context"
	balance2 "github.com/Iglesys347/equity/api/handlers/balance"
	"github.com/Iglesys347/equity/api/handlers/expenses"
	"github.com/Iglesys347/equity/api/handlers/users"
	"github.com/Iglesys347/equity/db"
	"github.com/Iglesys347/equity/logger"
	"github.com/Iglesys347/equity/recurring"
	"github.com/rs/zerolog/hlog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func requestLogger(next http.Handler) http.Handler {
	l := logger.Get()

	h := hlog.NewHandler(l)

	accessHandler := hlog.AccessHandler(
		func(r *http.Request, status, size int, duration time.Duration) {
			r = r.WithContext(l.WithContext(r.Context()))
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status_code", status).
				Int("response_size_bytes", size).
				Dur("elapsed_ms", duration).
				Msg("incoming request")
		},
	)

	userAgentHandler := hlog.UserAgentHandler("http_user_agent")

	return h(accessHandler(userAgentHandler(next)))
}

func main() {
	l := logger.Get()

	// Create application context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to the database
	database, err := db.Connect()
	defer database.Close()
	if err != nil {
		l.Fatal().Err(err).Msg("Unable to connect to DB")
	}

	// Create the recurrent expense renewer (run every 24 hours)
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				l.Info().Msg("renewing recurring expenses")
				err := recurring.CheckRecurring()
				if err != nil {
					l.Error().Err(err).Msg("failed to renew recurring expenses")
				} else {
					l.Info().Msg("successfully renewed recurring expenses")
				}
			case <-ctx.Done():
				break
			}
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Create a new HTTP server
	mux := http.NewServeMux()

	// Handling GET and POST requests on /expenses
	mux.HandleFunc("/expenses", expenses.ExpensesHandler)
	// Handling GET and PUT requests on /expenses/{id}
	mux.HandleFunc("/expenses/", expenses.ExpenseWithIDHandler)

	// Handling GET and POST requests on /users
	mux.HandleFunc("/users", users.UsersHandler)
	// Handling GET and PUT requests on /users/{id}
	mux.HandleFunc("/users/", users.UserWithIDHandler)

	mux.HandleFunc("/balance/", balance2.BalanceHandler)

	server := http.Server{
		Addr:    ":" + port,
		Handler: requestLogger(mux),
	}

	// Start the server
	l.Info().Str("port", port).Msgf("Starting Equity server on port %s", port)

	if err := server.ListenAndServe(); err != nil {
		l.Fatal().Err(err).Msg("Unable to start server")
	}

	// Wait for SIGINT.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
