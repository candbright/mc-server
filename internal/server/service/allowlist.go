package service

import (
	"github.com/candbright/server-mc/internal/errors"
	"github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/internal/server/manager"
)

type AllowListService struct {
	Manager *manager.Manager
}

func (service *AllowListService) List() model.AllowList {
	return service.Manager.AllowListManager.Data
}

func (service *AllowListService) Get(xuid string) (model.AllowListUser, error) {
	for _, user := range service.Manager.AllowListManager.Data {
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
