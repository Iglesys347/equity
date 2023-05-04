package balance

import (
	"github.com/Iglesys347/equity/models"
	"math"
	"sort"
)

func BalanceExpenses(users models.BalanceUsers) []models.Transfer {
	totalExpenses := 0.0
	for _, user := range users {
		totalExpenses += user.Expenses
	}

	for i := range users {
		users[i].Balance = (totalExpenses * users[i].Ratio) - users[i].Expenses
	}

	// BalanceUsers are sorted by Balance
	sort.Sort(users)

	transfers := []models.Transfer{}

	i, j := 0, len(users)-1
	for i < j {
		if users[i].Balance == 0 {
			i++
		} else if users[j].Balance == 0 {
			j--
		} else {
			transferAmount := math.Min(math.Abs(users[i].Balance), math.Abs(users[j].Balance))
			transfers = append(transfers, models.Transfer{
				From:   users[j].Name,
				To:     users[i].Name,
				Amount: transferAmount,
			})

			users[j].Balance -= transferAmount
			users[i].Balance += transferAmount
			if users[i].Balance == 0 {
				i++
			}
			if users[j].Balance == 0 {
				j--
			}
		}
	}

	return transfers
}
