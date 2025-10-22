package storage

import (
	"errors"
	"net/url"
	"skeleton/conf"
	"skeleton/model"
	"skeleton/storage/postgres"
)

type Storage interface {
	ClientCount(params url.Values) (total int, err error)
	ClientRead(params url.Values) (users []model.Client, total int, err error)
}

func NewStorage(databaseType string, cfg conf.DbCfg) (Storage, error) {
	switch databaseType {
	case "postgres":
		return postgres.NewPostgresStorage(cfg)
	default:
		return nil, errors.New("unknown storage type")
	}
}
