package core

type DomainError struct {
	Message string
}

func (self DomainError) Error() string {
	return self.Message
}
