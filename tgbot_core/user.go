package tgbot_core

type User struct {
	Id        int    `json:"id"`
	isBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Language  string `json:language_code`
}

func (u *User) getFullName() string {
	return u.FirstName + " " + u.LastName
}
