package process

import (
	model2 "github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"path"
)

const (
	AllowListFile = "allowlist.json"
)

type AllowList struct {
	AllowListFile *fm.FileManager[model2.AllowList]
}

func NewAllowList(dir string) *AllowList {
	fileManager := fm.Default[model2.AllowList](path.Join(dir, AllowListFile))
	return &AllowList{
		AllowListFile: fileManager,
	}
}
