package main

import (
	"regexp"
)

var rxEmail = regexp.MustCompile("^.+@.+\\..+$")
var rxName = regexp.MustCompile("^[a-zA-Z-]+$")

type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

type Member struct {
	Name   string
	Email  string
	Date   string
	Errors map[string]string
}

func (member *Member) Validate() bool {
	member.Errors = make(map[string]string)

	if !rxName.Match([]byte(member.Name)) {
		member.Errors["Name"] = "Please enter a valid name"
	}

	if !rxEmail.Match([]byte(member.Email)) {
		member.Errors["Email"] = "Please enter a valid email address"
	}

	return len(member.Errors) == 0
}
