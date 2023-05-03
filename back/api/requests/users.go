package requests

type NewUser struct {
	Name  string  `json:"name"`
	Wage  float64 `json:"wage,omitempty"`
	Ratio float32 `json:"ratio,omitempty"`
}

func (u *NewUser) Valid() bool {
	if u.Name == "" {
		return false
	}
	return true
}

type UpdateUser struct {
	Name  string  `json:"name,omitempty"`
	Wage  float64 `json:"wage,omitempty"`
	Ratio float32 `json:"ratio,omitempty"`
}

func (u *UpdateUser) Valid() bool {
	if u.Name == "" && u.Wage == 0.0 && u.Ratio == 0.0 {
		return false
	}
	return true
}
