package postgres

import (
	"fmt"
	"skeleton/conf"

	dbr "github.com/mailru/dbr"

	_ "github.com/lib/pq"
)

type DbCfg struct {
}

type Postgres struct {
	connection *dbr.Connection
	cfg        conf.DbCfg
}

func NewPostgresStorage(cfg conf.DbCfg) (*Postgres, error) {
	p := Postgres{
		cfg: cfg,
	}
	var err error
	p.connection, err = dbr.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		p.cfg.User,
		p.cfg.Password,
		p.cfg.Host,
		p.cfg.Port,
		p.cfg.DB,
	), nil)

	return &p, err
}
