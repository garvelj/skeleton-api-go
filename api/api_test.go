package api

import (
	"skeleton/conf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresConn(t *testing.T) {
	mockCfg := conf.DbCfg{
		Host:     "localhost",
		Port:     13131,
		User:     "nikola",
		Password: "password",
		DB:       "devDb",
	}

	db := PostgressConn(mockCfg)
	assert.NotNil(t, db)
}
