package service

import (
	"github.com/candbright/server-mc/internal/errors"
	"github.com/candbright/server-mc/internal/manager/process"
	"github.com/candbright/server-mc/internal/model"
)

type AllowListService struct {
	Manager *process.Process
}

func (service *AllowListService) List() model.AllowList {
	return service.Manager.AllowListFile.Data
}

func (service *AllowListService) Get(xuid string) (model.AllowListUser, error) {
	for _, user := range service.Manager.AllowListFile.Data {
		if user.XUid == xuid {
			return user, nil
		}
	}
	return model.AllowListUser{}, errors.NotExistError("AllowListUser", xuid)
}

func (service *AllowListService) Add(name string) error {
	panic("todo")
}

func (service *AllowListService) Remove(name string) error {
	panic("todo")
}
