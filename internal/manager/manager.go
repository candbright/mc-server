package manager

import (
	"github.com/candbright/server-mc/internal/manager/process"
)

type Config struct {
	Version string
	RootDir string
}

type Manager struct {
	*DownloadManager
	Server *process.Process
}

func New(cfg *Config) *Manager {
	manager := &Manager{
		DownloadManager: &DownloadManager{
			RootDir: cfg.RootDir,
			Version: cfg.Version,
		},
		Server: process.New(&process.Config{
			RootDir:        cfg.RootDir,
			AllowListFile:  "allowlist.json",
			PropertiesFile: "server.properties",
		}),
	}
	return manager
}

func (m *Manager) Upgrade() error {
	//1. 获取最新版本
	oldDownload := m.DownloadManager
	newVersion := ""
	//2. 若最新版本和当前版本不同，则下载最新版本
	if newVersion != oldDownload.Version {
		newDownload := &DownloadManager{
			RootDir: m.RootDir,
			Version: newVersion,
		}
		err := newDownload.Download()
		if err != nil {
			return err
		}
		m.DownloadManager = newDownload
	}
	//3. 复制旧版本文件到新版本
	return nil
}
