package core

import (
	"project/connection"
	"project/domain"
	"project/domain/infra"
	"project/infra/config"
)

type Ctx struct {
	infra domain.Infra
	con   domain.Connection
}

type infr struct {
	cfg infra.Config
}

func (s *infr) Config() infra.Config {
	return s.cfg
}

func (c *Ctx) Infra() domain.Infra {
	return c.infra
}

func (c *Ctx) Connection() domain.Connection {
	return c.con
}

func (c *Ctx) Make() domain.Context {
	return &Ctx{
		infra: c.infra,
		con:   c.con,
	}
}

func InitCtx() *Ctx {
	cfg := config.Make()
	db, err := connection.Make(cfg)
	if err != nil {
		panic(err)
	}

	return &Ctx{
		infra: &infr{
			cfg: cfg,
		},
		con: db,
	}
}
