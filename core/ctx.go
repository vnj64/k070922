package core

import "project/domain"

type Ctx struct {
	infra domain.Infra
	con   domain.Connection
}

type infr struct {
}

func (c *Ctx) Infra() domain.Infra {
	return c.infra
}

func (c *Ctx) Connection() domain.Connection {
	return c.con
}

func (c *Ctx) Make() domain.Context {
	return &Ctx{}
}

func InitCtx() *Ctx {
	return &Ctx{}
}
