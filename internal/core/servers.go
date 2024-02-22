package core

import (
	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/go-ssh/ssh/options"
	"github.com/go-resty/resty/v2"
	"path"
)

type Config struct {
	LatestVersion string
	RootDir       string
	SSHOptions    []options.Option
}

func New(cfg *Config) *Servers {
	session, err := ssh.NewSession(cfg.SSHOptions...)
	if err != nil {
		panic(err)
	}
	latestServer, err := NewServer(&ServerConfig{
		Version: cfg.LatestVersion,
		RootDir: cfg.RootDir,
		Session: session,
	})
	if err != nil {
		panic(err)
	}
	servers := &Servers{
		Session: session,
		Latest:  latestServer,
		client:  resty.New(),
	}
	return servers
}

type Servers struct {
	Session  ssh.Session
	Latest   *Server
	Previous map[string]*Server
	client   *resty.Client
}

func (servers *Servers) LatestVersion() (string, error) {
	//TODO
	/*resp, err := m.client.R().
		Get("https://www.minecraft.net/en-us/download/server/bedrock")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil*/
	return "1.20.62.02", nil
}

func (servers *Servers) VersionScan() ([]string, error) {
	//TODO
	return nil, nil
}

func (servers *Servers) Upgrade() error {
	//1. 获取最新版本
	oldServer := servers.Latest
	newVersion, err := servers.LatestVersion()
	if err != nil {
		return err
	}
	//2. 若最新版本和当前版本不同，则下载最新版本
	if newVersion == oldServer.Version {
		return nil
	}
	newServer, err := NewServer(&ServerConfig{
		Version: newVersion,
		RootDir: oldServer.RootDir,
		Session: servers.Session,
	})
	if err != nil {
		return err
	}
	err = newServer.Download()
	if err != nil {
		return err
	}
	servers.Latest = newServer
	servers.Previous[oldServer.Version] = oldServer
	//3. 复制旧版本数据文件到新版本
	err = servers.Session.Run("cp", "-r",
		path.Join(oldServer.ServerDir(), "world"),
		path.Join(servers.Latest.ServerDir()+"/"))
	if err != nil {
		return err
	}
	err = servers.Session.Run("cp",
		path.Join(oldServer.ServerDir(), allowListFile),
		path.Join(servers.Latest.ServerDir(), allowListFile))
	if err != nil {
		return err
	}
	err = servers.Session.Run("cp",
		path.Join(oldServer.ServerDir(), propertiesFile),
		path.Join(servers.Latest.ServerDir(), propertiesFile))
	if err != nil {
		return err
	}
	return nil
}
