package db

import (
	"database/sql"
	"fmt"
	"github.com/Iglesys347/equity/models"
	"strings"
)

func InsertExpense(expense models.Expense) (int, error) {
	sqlCmd := "INSERT INTO expenses (date, category, amount, description, user_id, recurring_expense_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	result := db.QueryRow(sqlCmd, expense.Date, expense.Category, expense.Amount, expense.Description, expense.UserId, expense.RecurringId)
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

func GetAllExpenses() ([]models.Expense, error) {
	rows, err := db.Query("SELECT id, date, category, amount, description, user_id FROM expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	expenses := []models.Expense{}
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ID, &expense.Date, &expense.Category, &expense.Amount, &expense.Description, &expense.UserId)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func GetExpense(id int) (*models.Expense, error) {
	sqlCmd := "SELECT id, date, category, amount, description, user_id FROM expenses WHERE id=$1"
	result := db.QueryRow(sqlCmd, id)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var expense models.Expense
	err := result.Scan(&expense.ID, &expense.Date, &expense.Category, &expense.Amount, &expense.Description, &expense.UserId)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func GetLastRecurringExpense(recurringId int) (*models.Expense, error) {
	sqlCmd := "SELECT id, date, category, amount, description, user_id FROM expenses WHERE recurring_expense_id=$1 ORDER BY date DESC LIMIT 1"
	result := db.QueryRow(sqlCmd, recurringId)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var expense models.Expense
	err := result.Scan(&expense.ID, &expense.Date, &expense.Category, &expense.Amount, &expense.Description, &expense.UserId)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func GetUserExpense(userId int) (*models.Expense, error) {
	sqlCmd := "SELECT id, date, category, amount, description, user_id FROM expenses WHERE user_id=$1"
	result := db.QueryRow(sqlCmd, userId)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var expense models.Expense
	err := result.Scan(&expense.ID, &expense.Date, &expense.Category, &expense.Amount, &expense.Description, &expense.UserId)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func GetUserTotalExpense(userId int, year int, month int) (float64, error) {
	sqlCmd := "SELECT SUM(amount) FROM expenses WHERE user_id=$1 AND EXTRACT(YEAR FROM date) = $2 AND EXTRACT(MONTH FROM date) = $3"
	result := db.QueryRow(sqlCmd, userId, year, month)
	if result.Err() != nil {
		return 0, result.Err()
	}
	var expense float64
	err := result.Scan(&expense)
	if err != nil {
		return 0, err
	}
	return expense, nil
}

func UpdateExpense(expense models.Expense, overwrite bool) (bool, error) {
	var sqlStatement string
	var result sql.Result
	var err error
	if overwrite {
		sqlStatement = "UPDATE expenses SET date=$1, category=$2, amount=$3, description=$4, user_id=$5 WHERE id=$6"
		result, err = db.Exec(sqlStatement, expense.Date, expense.Category, expense.Amount, expense.Description, expense.UserId, expense.ID)
	} else {
		// Only updating not empty fields
		var fields []string
		var params []interface{}
		paramCounter := 1
		if !expense.Date.IsZero() {
			fields = append(fields, fmt.Sprintf("date=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.Date)
		}
		if expense.Category != "" {
			fields = append(fields, fmt.Sprintf("category=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.Category)
		}
		if expense.Amount != 0.0 {
			fields = append(fields, fmt.Sprintf("amount=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.Amount)
		}
		if expense.Description != "" {
			fields = append(fields, fmt.Sprintf("description=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.Description)
		}
		if expense.UserId != 0 {
			fields = append(fields, fmt.Sprintf("user_id=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.UserId)
		}
		params = append(params, expense.ID)
		sqlStatement = fmt.Sprintf("UPDATE expenses SET %s WHERE id=$%d", strings.Join(fields, ","), paramCounter)
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

func ExpenseExists(id int) bool {
	sqlCmd := "SELECT exists(SELECT 1 FROM expenses WHERE id=$1)"
	result := db.QueryRow(sqlCmd, id)
	if result.Err() != nil {
		return false
	}

	var expenseExists bool
	err := result.Scan(&expenseExists)
	if err != nil {
		return false
	}
	return expenseExists
}
