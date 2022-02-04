package main

import "github.com/x-auth/common/models"

type plugin struct {
	username string
	password string
}

func (p *plugin) Login(username string, password string) (models.Profile, bool) {
	if username == p.username && password == p.password {
		return models.Profile{
			Name:        "Foo Bar",
			FamilyName:  "Bar",
			GivenName:   "Foo",
			NickName:    "foobar",
			Email:       "foobar@example.com",
			PhoneNumber: "000000000",
		}, true
	}

	return models.Profile{}, false
}

func (p *plugin) Init(config map[string]string) {
	p.username = config["username"]
	p.password = config["password"]
}

var Plugin plugin
