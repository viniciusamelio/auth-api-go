package core

type AggregateRoot struct {
	AuthenticationService AuthenticationService
	SessionService        SessionService
	Session               Session
}
