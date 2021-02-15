package models

import "testing"

func TestUser(t *testing.T) {

	u := User{Name: "Test Name"}

	actual := u.GetName()
	expected := "Test Name"

	if expected != actual {
		t.Errorf("u.getName failed, expected %v, got %v", expected, actual)
	}
}
