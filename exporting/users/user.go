package users

// User struct
type User struct {
	Name  string
	Email string
}

// ChangeName method
func (u *User) ChangeName(name string) {
	u.Name = name
}

// ChangeEmail method
func (u *User) ChangeEmail(email string) {
	u.Email = email
}
