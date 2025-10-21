package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(confpath string) (*Cfg, error) {
	data, err := os.ReadFile(confpath)
	if err != nil {
		return nil, err
	}

	var cfg Cfg
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Cfg struct {
	Storages `yaml:"storages"`
}

type Storages struct {
	Postgres DbCfg `yaml:"postgres"`
}

type DbCfg struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
