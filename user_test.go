package main

import "testing"

func TestUser(t *testing.T) {

	u := User{name: "Test Name"}

	actual := u.getName()
	expected := "Test Name"

	if expected != actual {
		t.Errorf("u.getName failed, expected %v, got %v", expected, actual)
	}
}
