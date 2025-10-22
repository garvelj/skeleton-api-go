package postgres

import (
	"net/url"
	"skeleton/model"

	"github.com/mailru/dbr"
)

func (p Postgres) ClientCount(params url.Values) (total int, err error) {
	sess := p.connection.NewSession(nil)

	statement := sess.Select("COUNT(*)").From("client")
	for key, val := range params {
		statement.Where(dbr.Eq(key, val))
	}

	_, err = statement.Load(&total)
	return
}

func (p Postgres) ClientRead(params url.Values) (users []model.Client, total int, err error) {
	sess := p.connection.NewSession(nil)

	fields := params.Get("fields")
	if fields == "" {
		fields = "*"
	}
	params.Del("fields")

	total, err = p.ClientCount(params)
	if err != nil {
		return
	}

	statement := sess.Select(fields).From("client")
	for key, val := range params {
		statement.Where(dbr.Eq(key, val))
	}

	_, err = statement.Load(&users)

	return users, total, err
}
