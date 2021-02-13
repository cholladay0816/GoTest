package main

type User struct {
	name string
}

func newUser(name string) *User {
	u := User{name: name}
	return &u
}
func (u User) getName() string {
	return u.name
}
