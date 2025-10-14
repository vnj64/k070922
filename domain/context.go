package domain

type Context interface {
	Make() Context
	Connection() Connection
	Infra() Infra
}
