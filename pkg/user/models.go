package user

// User ... 
type User struct {
	ID       int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Country string `json:"country"`
}