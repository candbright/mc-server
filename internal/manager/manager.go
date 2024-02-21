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
	Session ssh.Session
	*VersionInfo
	*process.Process
	client *resty.Client
}

func New(cfg *Config) *Manager {
	session, err := ssh.NewSession(cfg.SSHOptions...)
	if err != nil {
		panic(err)
	}
	versionInfo := &VersionInfo{
		RootDir: cfg.RootDir,
		Version: cfg.Version,
	}
	p := process.New(&process.Config{
		RootDir: versionInfo.ServerDir(),
		Session: session,
	})
	manager := &Manager{
		Session:     session,
		VersionInfo: versionInfo,
		Process:     p,
		client:      resty.New(),
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

func (m *Manager) ZipExist() (bool, error) {
	//zip文件是否存在
	return m.Session.Exists(m.ZipFilePath())
}

func (m *Manager) ServerExist() (bool, error) {
	//服务器目录是否存在
	return m.Session.Exists(m.ServerDir())
}

func (m *Manager) Download() error {
	//1. 检测是否存在当前版本的服务器目录
	existS, err := m.ServerExist()
	if err != nil {
		return err
	}
	if existS {
		return nil
	}
	//2. 不存在当前版本的服务器目录，则检测是否存在当前版本的zip文件
	existZ, err := m.ZipExist()
	if err != nil {
		return err
	}
	//2. 不存在当前版本的zip文件，先下载
	if !existZ {
		err = m.Session.Run("wget", m.DownloadUrl(), "-P", m.RootDir)
		if err != nil {
			return err
		}
	}

	//3. 解压zip文件
	err = m.Session.Run("unzip", m.ZipFilePath(), "-d", m.RootDir)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Upgrade() error {
	//1. 获取最新版本
	oldVersionInfo := m.VersionInfo
	newVersion, err := m.LatestVersion()
	if err != nil {
		return err
	}
	//2. 若最新版本和当前版本不同，则下载最新版本
	if newVersion != oldVersionInfo.Version {
		m.VersionInfo = &VersionInfo{
			RootDir: m.RootDir,
			Version: newVersion,
		}
	}
	//3. 复制旧版本数据文件到新版本
	err = m.Session.Run("cp", "-r",
		path.Join(oldVersionInfo.ServerDir(), "world"),
		path.Join(m.VersionInfo.ServerDir()+"/"))
	if err != nil {
		return err
	}
	err = m.Session.Run("cp",
		path.Join(oldVersionInfo.ServerDir(), "allowlist.json"),
		path.Join(m.VersionInfo.ServerDir(), "allowlist.json"))
	if err != nil {
		return err
	}
	err = m.Session.Run("cp",
		path.Join(oldVersionInfo.ServerDir(), "server.properties"),
		path.Join(m.VersionInfo.ServerDir(), "server.properties"))
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Delete() error {
	//1. 删除服务器目录
	existS, err := m.ServerExist()
	if err != nil {
		return err
	}
	if existS {
		err = m.Session.RemoveAll(m.ServerDir())
		if err != nil {
			return err
		}
	}
	//2. 删除zip文件
	existZ, err := m.ZipExist()
	if err != nil {
		return err
	}
	if existZ {
		err = m.Session.Remove(m.ZipFilePath())
		if err != nil {
			return err
		}
	}
	return nil
}
