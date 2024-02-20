package manager

import (
	"github.com/candbright/server-mc/internal/manager/process"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type Config struct {
	Version string
	RootDir string
}

type Manager struct {
	*DownloadUtil
	*process.Process
	client *resty.Client
}

func New(cfg *Config) *Manager {
	download := &DownloadUtil{
		RootDir: cfg.RootDir,
		Version: cfg.Version,
	}
	p := process.New(&process.Config{
		RootDir: download.ServerDir(),
	})
	manager := &Manager{
		DownloadUtil: download,
		Process:      p,
		client:       resty.New(),
	}
	return manager
}

func (m *Manager) LatestVersion() (string, error) {
	resp, err := m.client.R().
		Get("https://www.minecraft.net/en-us/download/server/bedrock")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil
}

func (m *Manager) Upgrade() error {
	//1. 获取最新版本
	oldDownload := m.DownloadUtil
	newVersion, err := m.LatestVersion()
	if err != nil {
		return err
	}
	//2. 若最新版本和当前版本不同，则下载最新版本
	if newVersion != oldDownload.Version {
		newDownload := &DownloadUtil{
			RootDir: m.RootDir,
			Version: newVersion,
		}
		err = newDownload.Download()
		if err != nil {
			return err
		}
		m.DownloadUtil = newDownload
	}
	//3. 复制旧版本数据文件到新版本
	//TODO
	return nil
}
