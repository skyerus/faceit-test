package user

import "encoding/json"

// User ...
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

// MarshalJSON - override in order to remove password from payload
func (u User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		Country   string `json:"country"`
	}
	tmp.ID = u.ID
	tmp.FirstName = u.FirstName
	tmp.LastName = u.LastName
	tmp.Nickname = u.Nickname
	tmp.Email = u.Email
	tmp.Country = u.Country
	return json.Marshal(&tmp)
}

// Filter - search filter
type Filter map[string]string
