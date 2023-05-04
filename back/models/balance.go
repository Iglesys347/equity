package models

type BalanceUser struct {
	Name     string  `json:"name"`
	Expenses float64 `json:"expenses"`
	Ratio    float64 `json:"ratio"`
	Balance  float64 `json:"balance"`
}

// BalanceUsers redefines interfaces to sort by BalanceUser.Balance
type BalanceUsers []BalanceUser

func (u BalanceUsers) Len() int {
	return len(u)
}
func (u BalanceUsers) Less(i, j int) bool {
	return u[i].Balance < u[j].Balance
}
func (u BalanceUsers) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u User) ToBalanceUser() BalanceUser {
	return BalanceUser{
		Name:     u.Name,
		Expenses: 0,
		Ratio:    float64(u.Ratio),
		Balance:  0,
	}
}

type Transfer struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type BalanceSummary struct {
	TotalExpenses float64      `json:"total-expenses"`
	Users         BalanceUsers `json:"users"`
	Balances      []Transfer   `json:"balances"`
}
