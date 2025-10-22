package model

import "time"

type Client struct {
	Id         string    `db:"id" json:"id,omitempty" yaml:"id,omitempty"`
	FirstName  string    `db:"first_name" json:"first_name,omitempty" yaml:"first_name,omitempty"`
	LastName   string    `db:"last_name" json:"last_name,omitempty" yaml:"last_name,omitempty"`
	Email      string    `db:"email" json:"email,omitempty" yaml:"email,omitempty"`
	Password   string    `db:"password" json:"password,omitempty" yaml:"password,omitempty"`
	CreateTime time.Time `db:"create_time" json:"create_time,omitempty" yaml:"create_time,omitempty"`
	UpdateTime time.Time `db:"last_time" json:"last_time,omitempty" yaml:"last_time,omitempty"`
}
