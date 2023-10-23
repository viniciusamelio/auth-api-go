package core

type DefaultError struct {
	Message string
	Code    int
}

func (self DefaultError) Error() string {
	return self.Message
}

func (self *DefaultError) SetMessage(message string) {
	self.Message = message
}
