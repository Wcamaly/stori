package user

type UserData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
}
