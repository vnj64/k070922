package domain

import "project/domain/infra"

type Infra interface {
	Config() infra.Config
}
