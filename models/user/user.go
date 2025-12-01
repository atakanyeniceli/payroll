package user

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUser() *User {
	return &User{
		ID:        0,
		Firstname: "",
		Lastname:  "",
		Email:     "",
		Password:  "",
	}
}
