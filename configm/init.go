package configm

import "sync"

type Item struct {
	Name         string `json:"name"`
	SSH_Host     string `json:"ssh_host"`
	SSH_Port     string `json:"ssh_port"`
	SSH_User     string `json:"ssh_user"`
	SSH_Password string `json:"ssh_password"`
}

type Manager struct {
	mu    sync.RWMutex
	items map[string]Item
}

var manager *Manager = &Manager{
	items: map[string]Item{},
}

func GetManager() *Manager {
	return manager
}
