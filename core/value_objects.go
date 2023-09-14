package core

import "regexp"

type Uuid struct {
	value string
}

func (self Uuid) set(value string) error {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	if !r.MatchString(value) {
		return DomainError{
			Message: "Invalid UUID",
		}
	}
	self.value = value
	return nil
}
