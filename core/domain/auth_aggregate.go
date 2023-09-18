package domain

type AuthAggregate struct {
	AuthenticationService AuthenticationService
	SessionService        SessionService
	Session               Session
}
