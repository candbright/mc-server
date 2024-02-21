package process

import (
	"github.com/candbright/server-mc/internal/errors"
	"github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"path"
)

const (
	AllowListFile = "allowlist.json"
)

type AllowList struct {
	AllowListFile *fm.FileManager[model.AllowList]
}

func NewAllowList(dir string) *AllowList {
	fileManager := fm.Default[model.AllowList](path.Join(dir, AllowListFile))
	return &AllowList{
		AllowListFile: fileManager,
	}
}

func (fm *AllowList) AddUser(user model.AllowListUser) error {
	fm.AllowListFile.Data = append(fm.AllowListFile.Data, user)
	err := fm.AllowListFile.Write()
	if err != nil {
		return err
	}
	return nil
}

func (fm *AllowList) DeleteUserById(id string) error {
	for i, user := range fm.AllowListFile.Data {
		if user.XUid == id {
			fm.AllowListFile.Data = append(fm.AllowListFile.Data[:i], fm.AllowListFile.Data[i+1:]...)
		}
	}
	return errors.NotExistError("user", id)
}

func (fm *AllowList) DeleteUserByName(name string) error {
	for i, user := range fm.AllowListFile.Data {
		if user.Name == name {
			fm.AllowListFile.Data = append(fm.AllowListFile.Data[:i], fm.AllowListFile.Data[i+1:]...)
		}
	}
	return errors.NotExistError("user", name)
}
