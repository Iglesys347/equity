package db

import (
	"database/sql"
	"fmt"
	"github.com/Iglesys347/equity/models"
	"strings"
)

func InsertUser(user models.User) (int, error) {
	sqlCmd := "INSERT INTO users (name, wage, ratio) VALUES ($1, $2, $3) RETURNING id"
	result := db.QueryRow(sqlCmd, user.Name, user.Wage, user.Ratio)
	id := -1
	if result.Err() != nil {
		return id, result.Err()
	}
	err := result.Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func GetAllUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT id, name, wage, ratio FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Wage, &user.Ratio)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id int) (*models.User, error) {
	sqlCmd := "SELECT id, name, wage, ratio FROM users WHERE id=$1"
	result := db.QueryRow(sqlCmd, id)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user models.User
	err := result.Scan(&user.ID, &user.Name, &user.Wage, &user.Ratio)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user models.User, overwrite bool) (bool, error) {
	var sqlStatement string
	var result sql.Result
	var err error
	if overwrite {
		sqlStatement = "UPDATE users SET name=$1, wage=$2, ratio=$3 WHERE id=$5"
		result, err = db.Exec(sqlStatement, user.Name, user.Wage, user.Ratio, user.ID)
	} else {
		// Only updating not empty fields
		var fields []string
		var params []interface{}
		paramCounter := 1
		if user.Name != "" {
			fields = append(fields, fmt.Sprintf("name=$%d", paramCounter))
			paramCounter++
			params = append(params, user.Name)
		}
		if user.Wage != 0.0 {
			fields = append(fields, fmt.Sprintf("wage=$%d", paramCounter))
			paramCounter++
			params = append(params, user.Wage)
		}
		if user.Ratio != 0.0 {
			fields = append(fields, fmt.Sprintf("ratio=$%d", paramCounter))
			paramCounter++
			params = append(params, user.Ratio)
		}
		params = append(params, user.ID)
		sqlStatement = fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", strings.Join(fields, ","), paramCounter)
		result, err = db.Exec(sqlStatement, params...)
	}
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected == 1, nil
}

func UserExists(id int) bool {
	sqlCmd := "SELECT exists(SELECT 1 FROM users WHERE id=$1)"
	result := db.QueryRow(sqlCmd, id)
	if result.Err() != nil {
		return false
	}

	var userExists bool
	err := result.Scan(&userExists)
	if err != nil {
		return false
	}
	return userExists
}
