package models

// User ...
type User struct {
	Name string
}

// GetName fetches name ...
func (u User) GetName() string {
	return u.Name
}
