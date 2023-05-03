package requests

import "time"

type NewExpense struct {
	Date        time.Time `json:"date"`
	Category    string    `json:"category,omitempty"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description,omitempty"`
}

func (e *NewExpense) Valid() bool {
	if e.Date.IsZero() {
		return false
	}
	if e.Amount == 0.0 {
		return false
	}
	return true
}

type UpdateExpense struct {
	Date        time.Time `json:"date,omitempty"`
	Category    string    `json:"category,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (e *UpdateExpense) Valid() bool {
	if e.Date.IsZero() && e.Category == "" && e.Amount == 0.0 && e.Description == "" {
		return false
	}
	return true
}
