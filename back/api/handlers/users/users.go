package users

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

func UsersHandler(w http.ResponseWriter, r *http.Request) {
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
		l.Info().Msgf("incoming GET request on /users with search query '%s' on page '%s'", searchQuery, pageNum)

		// Retrieve all users
		users, err := db.GetAllUsers()
		if err != nil {
			l.Error().Err(err).Msg("unable to retrieve users")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		l.Debug().Interface("get_users_response", users).Send()

		// Encode users as JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

		l.Trace().Msgf("search query '%s' succeeded without errors", searchQuery)

	case "POST":
		// Decode the request body into a new user
		var newUser requests.NewUser
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			l.Error().Err(err).Msg("failed to decode JSON body")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		l.Debug().Interface("post_user_body", newUser).Send()

		// Checking the user
		if !newUser.Valid() {
			l.Error().Interface("post_user_body", newUser).Msg("user is not valid")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		// Insert the new user into the database
		db.InsertUser(models.User{
			Name:  newUser.Name,
			Wage:  newUser.Wage,
			Ratio: newUser.Ratio,
		})

		// Send a success response
		w.WriteHeader(http.StatusCreated)

		l.Trace().Msg("post new user succeeded without errors")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserWithIDHandler(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())

	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/users/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		l.Error().Err(err).Msg("invalid ID")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("method", r.Method).Int("user_id", id)
	})
	switch r.Method {
	case "GET":
		l.Info().Msgf("incoming GET request on /users/ with user ID '%s'", id)

		// Retrieve the user
		user, err := db.GetUser(id)
		if err != nil {
			l.Error().Err(err).Msg("unable to retrieve user")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		l.Debug().Interface("get_user_response", user).Send()

		// Encode user as JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

		l.Trace().Msgf("get user with ID %s succeeded without error", id)

	case "PUT":
		// Decode the request body into a new user
		var updateUser requests.UpdateUser
		err := json.NewDecoder(r.Body).Decode(&updateUser)
		if err != nil {
			l.Error().Err(err).Msg("failed to decode JSON body")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		l.Debug().Interface("put_user_body", updateUser).Send()

		// Checking the user
		if !updateUser.Valid() {
			l.Error().Interface("post_user_body", updateUser).Msg("user is not valid")
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		// Checking if the user with given ID exists
		if !db.UserExists(id) {
			l.Error().Msgf("user with ID %d does not exists", id)
			http.Error(w, fmt.Sprintf("User with ID %d does not exists", id), http.StatusBadRequest)
			return
		}

		// Update the user
		updated, err := db.UpdateUser(models.User{
			ID:    id,
			Name:  updateUser.Name,
			Wage:  updateUser.Wage,
			Ratio: updateUser.Ratio,
		}, false)
		if err != nil {
			l.Error().Err(err).Msg("error while updating user")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		if !updated {
			l.Error().Msg("unable to update user")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)

		l.Trace().Msgf("put user with ID %s succeeded without errors", id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
