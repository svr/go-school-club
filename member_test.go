package main

import (
	"strconv"
	"testing"
)

var testData = []struct {
	description string
	member      Member
	want        bool
}{
	{"Valid Name and Email", Member{"Name", "test@example.com", "01.01.2001", nil}, true},
	{"Invalid Name", Member{"__Invalid name", "test@example.com", "01.01.2001", nil}, false},
	{"Invalid Email", Member{"Name", "invalid.email@", "01.01.2001", nil}, false},
	{"Invalid Name and Email", Member{"111", "--==", "01.01.2001", nil}, false},
	{"Empty Name and Email", Member{"", "", "01.01.2001", nil}, false},
}

func TestValidate(t *testing.T) {
	for _, dataItem := range testData {
		t.Run(dataItem.description, func(t *testing.T) {
			actual := dataItem.member.Validate()
			if actual != dataItem.want {
				t.Error("Expected validation to be '" + strconv.FormatBool(dataItem.want) + "' for " + dataItem.member.Name + "/" + dataItem.member.Email)
			}
		})

	}
}

func TestInvalidNameMessage(t *testing.T) {
	member := Member{"__Invalid name", "test@example.com", "01.01.2001", nil}
	member.Validate()
	actual := member.Errors["Name"]

	assertString(t, actual, "Please enter a valid name")
}

func TestInvalidEmailMessage(t *testing.T) {
	member := Member{"Name", "@@@", "01.01.2001", nil}
	member.Validate()
	actual := member.Errors["Email"]

	assertString(t, actual, "Please enter a valid email address")
}

func assertString(t *testing.T, actual string, want string) {
	if actual != want {
		t.Error("Failed: expected '" + want + "' recieved '" + actual + "'")
	}
}
