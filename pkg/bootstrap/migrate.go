package bootstrap

import (
	migration "ocserv/internal/database"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	migration.Migrate()
}
