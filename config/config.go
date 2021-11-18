package config

import (
	"encoding/json"
	"io/ioutil"
)

type Dbconfig struct {
	DB_USERNAME string `json:"DB_USERNAME"`
	DB_PASSWORD string `json:"DB_PASSWORD"`
	DB_HOST     string `json:"DB_HOST"`
	DB_NAME     string `json:"DB_NAME"`
}

func Getconfig() Dbconfig {

	conf := Dbconfig{}

	cred, err := ioutil.ReadFile("config.json")
	if err == nil {
		_ = json.Unmarshal([]byte(cred), &conf)
	}

	return conf
}
