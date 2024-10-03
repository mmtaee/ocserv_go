package ocserv

import (
	"ocserv/pkg/config"
)

type CMDRepository struct{}

type HookRepository struct{}

type ServiceOcservRepositoryInterface interface {
	ReloadService() error
	CreateOrUpdateUser(string, string, string) error
	ChangeGroup(string, string) error
	Lock(string) error
	Unlock(string) error
	DeleteUser(string) error
	Disconnect(string) error
	OnlineUsers(bool) (interface{}, error)
	SyncUsers() ([][2]string, error)
	ShowIPBans() interface{}
	ShowIPBansPoints() interface{}
	UnBanIP(string) error
	ShowStatus() (string, error)
	ShowIRoutes() ([]map[string]interface{}, error)
}

func NewOcservRepository() ServiceOcservRepositoryInterface {
	cfg := config.GetApp()
	if cfg.Hook {
		return &HookRepository{}
	}
	return &CMDRepository{}
}
