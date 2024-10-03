package ocserv

func (h *HookRepository) ReloadService() error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) CreateOrUpdateUser(group, username, password string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) ChangeGroup(s string, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) Lock(s string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) Unlock(s string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) DeleteUser(username string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) Disconnect(username string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) OnlineUsers(onlyUsername bool) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) SyncUsers() ([][2]string, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) ShowIPBans() interface{} {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) ShowIPBansPoints() interface{} {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) UnBanIP(ip string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) ShowStatus() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HookRepository) ShowIRoutes() ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}
