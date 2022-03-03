package main

import (
	"github.com/x-auth/common/models"
	"github.com/x-auth/common/plugins"
)

type Plugin struct {
	username string
	password string
}

func (p *Plugin) Login(username string, password string) (models.Profile, bool) {
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

func NewPlugin(cfg map[string]string) plugins.AuthPlugin {
	return &Plugin{
		username: cfg["username"],
		password: cfg["password"],
	}
}
