package manager

import (
	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/go-ssh/ssh/options"
	"github.com/candbright/server-mc/internal/manager/process"
	"github.com/go-resty/resty/v2"
	"path"
)

type Config struct {
	Version    string
	RootDir    string
	SSHOptions []options.Option
}

type Manager struct {
	*DownloadUtil
	*process.Process
	client *resty.Client
}

func New(cfg *Config) *Manager {
	session, err := ssh.NewSession(cfg.SSHOptions...)
	if err != nil {
		panic(err)
	}
	download := &DownloadUtil{
		RootDir: cfg.RootDir,
		Version: cfg.Version,
		Session: session,
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
	/*resp, err := m.client.R().
		Get("https://www.minecraft.net/en-us/download/server/bedrock")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil*/
	return "1.20.62.02", nil
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
	err = m.DownloadUtil.Session.Run("cp", "-r",
		path.Join(oldDownload.ServerDir(), "world"),
		path.Join(m.DownloadUtil.ServerDir()+"/"))
	if err != nil {
		return err
	}
	err = m.DownloadUtil.Session.Run("cp",
		path.Join(oldDownload.ServerDir(), "allowlist.json"),
		path.Join(m.DownloadUtil.ServerDir(), "allowlist.json"))
	if err != nil {
		return err
	}
	err = m.DownloadUtil.Session.Run("cp",
		path.Join(oldDownload.ServerDir(), "server.properties"),
		path.Join(m.DownloadUtil.ServerDir(), "server.properties"))
	if err != nil {
		return err
	}
	return nil
}
