package balance

import (
	"encoding/json"
	"fmt"
	"github.com/Iglesys347/equity/balance"
	"github.com/Iglesys347/equity/db"
	"github.com/Iglesys347/equity/logger"
	"github.com/Iglesys347/equity/models"
	"github.com/rs/zerolog"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var l = logger.Get()

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())

	path := r.URL.Path
	params := strings.TrimPrefix(path, "/balance/")
	match, _ := regexp.MatchString(`^\d{4}\/(0?[1-9]|1[012])$`, params)
	fmt.Println(params, match)
	if !match {
		l.Error().Str("param", params).Msg("invalid URI")
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}
	paramsSplit := strings.Split(params, "/")
	year, err := strconv.Atoi(paramsSplit[0])
	if err != nil {
		l.Error().Str("param", params).Msg("invalid year")
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}
	month, err := strconv.Atoi(paramsSplit[1])
	if err != nil {
		l.Error().Str("param", params).Msg("invalid month")
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return
	}

	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("method", r.Method).Int("year", year).Int("month", month)
	})
	switch r.Method {
	case "GET":
		l.Info().Msgf("incoming GET request on /balance/%d/%d", year, month)

		// Retrieve the users
		users, err := db.GetAllUsers()
		if err != nil {
			l.Error().Err(err).Msg("unable to retrieve users")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		// Retrieve the total expenses for all users
		balanceUsers := make(models.BalanceUsers, len(users))
		totalExpenses := 0.0
		for i, usr := range users {
			balanceUsers[i] = usr.ToBalanceUser()
			// Retrieve total expenses for the month & year
			usrTotalExpenses, err := db.GetUserTotalExpense(usr.ID, year, month)
			if err != nil {
				l.Error().Err(err).Msg("unable to retrieve user expenses")
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}
			balanceUsers[i].Expenses = usrTotalExpenses
			totalExpenses += usrTotalExpenses
		}

		// Compute the transfers
		transfers := balance.BalanceExpenses(balanceUsers)

		result := models.BalanceSummary{
			TotalExpenses: totalExpenses,
			Users:         balanceUsers,
			Balances:      transfers,
		}

		// Encode result as JSON and send it
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

		l.Trace().Msgf("get balance with year=%d and month=%d succeeded without error", year, month)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
