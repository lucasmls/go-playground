package domain

// User ..
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
}
