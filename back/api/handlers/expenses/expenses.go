package expenses

import (
	"encoding/json"
	"fmt"
	"github.com/Iglesys347/equity/api/requests"
	"github.com/Iglesys347/equity/db"
	"github.com/Iglesys347/equity/logger"
	"github.com/Iglesys347/equity/models"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
	"strings"
)

var l = logger.Get()

func ExpensesHandler(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())

	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("method", r.Method)
	})
	switch r.Method {
	case "GET":
		// Search and pagination params
		params := r.URL.Query()
		searchQuery := params.Get("q")
		pageNum := params.Get("page")
		if pageNum == "" {
			pageNum = "1"
		}

		l.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("search_query", searchQuery).Str("page_num", pageNum)
		})
		l.Info().Msgf("incoming GET request on /expenses with search query '%s' on page '%s'", searchQuery, pageNum)

		// Retrieve all expenses
		expenses, err := db.GetAllExpenses()
		if err != nil {
			l.Error().Err(err).Msg("unable to retrieve expenses")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		l.Debug().Interface("get_expense_response", expenses).Send()

		// Encode expenses as JSON and send the response
		w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		json.NewEncoder(w).Encode(expenses)

		l.Trace().Msgf("search query '%s' succeeded without errors", searchQuery)

	case "POST":
		// Decode the request body into a new expense
		var newExpense requests.NewExpense
		err := json.NewDecoder(r.Body).Decode(&newExpense)
		if err != nil {
			l.Error().Err(err).Msg("failed to decode JSON body")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		l.Debug().Interface("post_expense_body", newExpense).Send()

		// Checking the expense
		if !newExpense.Valid() {
			l.Error().Interface("post_expense_body", newExpense).Msg("expense is not valid")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		// Insert the new expense into the database
		db.InsertExpense(models.Expense{
			Date:        newExpense.Date,
			Category:    newExpense.Category,
			Amount:      newExpense.Amount,
			Description: newExpense.Description,
		})

		// Send a success response
		w.WriteHeader(http.StatusCreated)

		l.Trace().Msg("post new expense succeeded without errors")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ExpenseWithIDHandler(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())

	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/expenses/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		l.Error().Err(err).Msg("invalid ID")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("method", r.Method).Int("expense_id", id)
	})
	switch r.Method {
	case "GET":
		l.Info().Msgf("incoming GET request on /expenses/ with expense ID '%s'", id)

		// Retrieve the expense
		expenses, err := db.GetExpense(id)
		if err != nil {
			l.Error().Err(err).Msg("unable to retrieve expense")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		l.Debug().Interface("get_expense_response", expenses).Send()

		// Encode expenses as JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expenses)

		l.Trace().Msgf("get expense with ID %s succeeded without error", id)

	case "PUT":
		// Decode the request body into a new expense
		var updateExpense requests.UpdateExpense
		err := json.NewDecoder(r.Body).Decode(&updateExpense)
		if err != nil {
			l.Error().Err(err).Msg("failed to decode JSON body")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		l.Debug().Interface("put_expense_body", updateExpense).Send()

		// Checking the expense
		if !updateExpense.Valid() {
			l.Error().Interface("post_expense_body", updateExpense).Msg("expense is not valid")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		// Checking if the expense with given ID exists
		if !db.ExpenseExists(id) {
			l.Error().Msgf("expense with ID %d does not exists", id)
			http.Error(w, fmt.Sprintf("Expense with ID %d does not exists", id), http.StatusBadRequest)
			return
		}

		// Update the expense
		updated, err := db.UpdateExpense(models.Expense{
			ID:          id,
			Date:        updateExpense.Date,
			Category:    updateExpense.Category,
			Amount:      updateExpense.Amount,
			Description: updateExpense.Description,
		}, false)
		if err != nil {
			l.Error().Err(err).Msg("error while updating expense")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		if !updated {
			l.Error().Msg("unable to update expense")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)

		l.Trace().Msgf("put expense with ID %s succeeded without errors", id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
