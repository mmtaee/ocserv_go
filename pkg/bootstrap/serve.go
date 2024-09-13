package bootstrap

import (
	"ocserv/pkg/config"
	"ocserv/pkg/database"
	"ocserv/pkg/routing"
)

func Serve() {
	config.LoadEnv()
	config.Set()
	database.Connect()
	routing.Serve()
}
