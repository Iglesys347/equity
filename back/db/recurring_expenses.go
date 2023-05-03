package db

import (
	"database/sql"
	"fmt"
	"github.com/Iglesys347/equity/models"
	"strings"
	"time"
)

func InsertRecurringExpense(expense models.RecurringExpense) (int, error) {
	sqlCmd := "INSERT INTO recurring_expenses (category, amount, description, frequency, start_date, end_date, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	result := db.QueryRow(sqlCmd, expense.Category, expense.Amount, expense.Description, expense.Frequency, expense.StartDate, expense.EndDate, expense.UserID)
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

func GetAllRecurringExpenses() ([]models.RecurringExpense, error) {
	rows, err := db.Query("SELECT id, category, amount, description, frequency, start_date, end_date, user_id FROM recurring_expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.RecurringExpense
	for rows.Next() {
		var expense models.RecurringExpense
		var endDate sql.NullTime
		var desc sql.NullString
		err := rows.Scan(&expense.Id, &expense.Category, &expense.Amount, &desc, &expense.Frequency, &expense.StartDate, &endDate, &expense.UserID)
		if endDate.Valid {
			expense.EndDate = endDate.Time
		} else {
			// setting the maximum possible date if end date is null
			expense.EndDate = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
		}
		if desc.Valid {
			expense.Description = desc.String
		}
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

func GetRecurringExpense(id int) (*models.RecurringExpense, error) {
	sqlCmd := "SELECT id, category, amount, description, frequency, start_date, end_date, user_id FROM recurring_expenses WHERE id=$1"
	result := db.QueryRow(sqlCmd, id)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var expense models.RecurringExpense
	var endDate sql.NullTime
	var desc sql.NullString
	err := result.Scan(&expense.Id, &expense.Category, &expense.Amount, &desc, &expense.Frequency, &expense.StartDate, &endDate, &expense.UserID)
	if endDate.Valid {
		expense.EndDate = endDate.Time
	} else {
		// setting the maximum possible date if end date is null
		expense.EndDate = time.Unix(1<<63-62135596801, 999999999)
	}
	if desc.Valid {
		expense.Description = desc.String
	}
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func UpdateRecurringExpense(expense models.RecurringExpense, overwrite bool) (bool, error) {
	var sqlStatement string
	var result sql.Result
	var err error
	if overwrite {
		sqlStatement = "UPDATE recurring_expenses  SET category=$1, amount=$2, description=$3, frequency=$4, start_date=$5, end_date=$6, user_id=$7 WHERE id=$8"
		result, err = db.Exec(sqlStatement, expense.Category, expense.Amount, expense.Description, expense.Frequency, expense.StartDate, expense.EndDate, expense.UserID, expense.Id)
	} else {
		// Only updating not empty fields
		var fields []string
		var params []interface{}
		paramCounter := 1
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
		if expense.Frequency != "" {
			fields = append(fields, fmt.Sprintf("frequecy=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.Frequency)
		}
		if !expense.StartDate.IsZero() {
			fields = append(fields, fmt.Sprintf("start_date=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.StartDate)
		}
		if !expense.EndDate.IsZero() {
			fields = append(fields, fmt.Sprintf("end_date=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.EndDate)
		}
		if expense.UserID != 0 {
			fields = append(fields, fmt.Sprintf("user_id=$%d", paramCounter))
			paramCounter++
			params = append(params, expense.UserID)
		}
		params = append(params, expense.Id)
		sqlStatement = fmt.Sprintf("UPDATE recurring_expenses SET %s WHERE id=$%d", strings.Join(fields, ","), paramCounter)
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

func RecurringExpenseInExpenses(id int) (bool, error) {
	// first checking if the recurring expense exists in expenses table
	sqlCmd := "SELECT exists(SELECT 1 FROM expenses WHERE recurring_expense_id=$1)"
	result := db.QueryRow(sqlCmd, id)
	if result.Err() != nil {
		l.Debug().Str("function", "RecurringExpenseInExpenses").Msg("")
		return false, result.Err()
	}
	var expenseExists bool
	err := result.Scan(&expenseExists)
	if err != nil {
		return false, err
	}
	if !expenseExists {
		return false, nil
	}

	// if the recurring expense exist in expenses, we need to check if it needs to be re-issued
	now := time.Now()
	recExpense, err := GetRecurringExpense(id)
	if err != nil {
		return false, err
	}
	lastExpense, err := GetLastRecurringExpense(id)
	if err != nil {
		return false, err
	}
	if recExpense.Frequency == "daily" {
		return lastExpense.Date.Day() >= now.Day(), nil
	}
	if recExpense.Frequency == "weekly" {
		expenseYear, expenseWeek := lastExpense.Date.ISOWeek()
		nowYear, nowWeek := now.ISOWeek()
		return expenseWeek >= nowWeek && expenseYear >= nowYear, nil
	}
	if recExpense.Frequency == "monthly" {
		return lastExpense.Date.Month() >= now.Month(), nil
	}
	if recExpense.Frequency == "yearly" {
		return lastExpense.Date.Year() >= now.Year(), nil
	}
	return false, nil
}
