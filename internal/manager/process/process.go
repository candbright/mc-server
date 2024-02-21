package process

import (
	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/server-mc/internal/model"
)

type Config struct {
	RootDir string
	Session ssh.Session
}

type Process struct {
	Session          ssh.Session
	AllowList        *AllowList
	ServerProperties *ServerProperties
}

func New(cfg *Config) *Process {
	manager := &Process{
		AllowList:        NewAllowList(cfg.RootDir),
		ServerProperties: NewServerProperties(cfg.RootDir),
	}
	if cfg.Session != nil {
		manager.Session = cfg.Session
	} else {
		session, err := ssh.NewSession()
		if err != nil {
			panic(err)
		}
		manager.Session = session
	}
	return manager
}

func (p *Process) Start() error {
	return p.Session.Run("systemctl", "start", "mc-server")
}

func (p *Process) Stop() error {
	return p.Session.Run("systemctl", "stop", "mc-server")
}

func (p *Process) Restart() error {
	return p.Session.Run("systemctl", "restart", "mc-server")
}

func (p *Process) AddUser(user model.AllowListUser) error {
	err := p.AllowList.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (p *Process) DeleteUser(id string) error {
	err := p.AllowList.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil
}
