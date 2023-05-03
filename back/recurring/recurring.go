package recurring

import (
	"github.com/Iglesys347/equity/db"
	"github.com/Iglesys347/equity/logger"
	"github.com/Iglesys347/equity/models"
	"github.com/rs/zerolog"
	"time"

	_ "github.com/lib/pq"
)

func CheckRecurring() error {
	l := logger.Get()
	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("type", "recurring_check")
	})

	recurringExpenses, err := db.GetAllRecurringExpenses()
	if err != nil {
		l.Error().Err(err).Msg("unable to retrieve recurring expenses")
		return err
	}
	if len(recurringExpenses) == 0 {
		l.Info().Msg("no recurring expenses to renew")
	}

	for _, recExp := range recurringExpenses {
		l.Debug().Interface("recurrent_expense", recExp).Msg("renewing recurrent expense")
		if recExp.IsDue() {
			l.Debug().Interface("recurrent_expense", recExp).Msg("expense is due, continue renewal")
			alreadyRenewed, err := db.RecurringExpenseInExpenses(recExp.Id)
			if err != nil {
				l.Error().Err(err).Int("recurrent_expense_id", recExp.Id).Msg("unable to check if recurrent expense exists in expenses table")
			}

			if err == nil && !alreadyRenewed {
				l.Debug().Interface("recurrent_expense", recExp).Msg("expense not already renewed, continue renewal")
				_, err := db.InsertExpense(models.Expense{
					Date:        time.Now(),
					Category:    recExp.Category,
					Amount:      recExp.Amount,
					Description: recExp.Description,
					UserId:      recExp.UserID,
					RecurringId: recExp.Id,
				})
				if err != nil {
					l.Error().Err(err).Int("recurrent_expense_id", recExp.Id).Msg("unable to renew recurrent expense")
				}
				l.Info().Interface("recurrent_expense", recExp).Msg("renewed recurrent expense")
			} else {
				l.Debug().Interface("recurrent_expense", recExp).Msg("recurring expense already renewed")
			}
		} else {
			l.Debug().Interface("recurrent_expense", recExp).Msg("recurring expense is not due")
		}

		//select {
		//case <-ctx.Done():
		//	break
		//}
	}
	return nil
}
