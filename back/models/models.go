package models

import (
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	RecurringId int       `json:"recurring_id"`
}

func (e *Expense) Valid() bool {
	return true
}

type RecurringExpense struct {
	Id          int       `json:"id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Frequency   string    `json:"frequency"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UserID      int       `json:"user_id"`
}

func (e *RecurringExpense) Valid() bool {
	if e.EndDate.IsZero() {
		e.EndDate = time.Unix(1<<63-62135596801, 999999999)
	}
	return true
}

func (e *RecurringExpense) IsDue() bool {
	now := time.Now()

	if e.EndDate.Before(now) {
		return false
	}

	switch e.Frequency {
	case "daily":
		return true
	case "weekly":
		return now.Weekday() == e.StartDate.Weekday()
	case "monthly":
		return now.Day() == e.StartDate.Day()
	case "yearly":
		return now.Month() == e.StartDate.Month() && now.Day() == e.StartDate.Day()
	default:
		return false
	}
}

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Wage  float64 `json:"wage"`
	Ratio float32 `json:"ratio"`
}
