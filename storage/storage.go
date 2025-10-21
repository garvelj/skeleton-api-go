package storage

import (
	"errors"
	"net/url"
	"skeleton/conf"
	"skeleton/model"
	"skeleton/storage/postgres"
)

type Storage interface {
	UserCount(params url.Values) (total int, err error)
	UserRead(params url.Values) (users []model.User, total int, err error)
}

func NewStorage(databaseType string, cfg conf.DbCfg) (Storage, error) {
	switch databaseType {
	case "postgres":
		return postgres.NewPostgresStorage(cfg)
	default:
		return nil, errors.New("unknown storage type")
	}
}
