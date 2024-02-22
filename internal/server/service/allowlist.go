package service

import (
	"github.com/candbright/server-mc/internal/core"
	"github.com/candbright/server-mc/internal/repo"
)

type ServerInfoService struct {
	Servers *core.Servers
}

func (service *ServerInfoService) List() repo.AllowList {
	return nil
}
