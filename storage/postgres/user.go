package postgres

import (
	"net/url"
	"skeleton/model"

	"github.com/mailru/dbr"
)

func (p Postgres) UserCount(params url.Values) (total int, err error) {
	sess := p.connection.NewSession(nil)

	statement := sess.Select("COUNT(*)").From("user")
	for key, val := range params {
		statement.Where(dbr.Eq(key, val))
	}

	_, err = statement.Load(&total)
	return
}

func (p Postgres) UserRead(params url.Values) (users []model.User, total int, err error) {
	return nil, 0, nil
}
