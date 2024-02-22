package core

import (
	"fmt"
	"github.com/candbright/go-ssh/ssh"
)

type Process struct {
	Version string
	Session ssh.Session
}

func NewProcess(cfg *ServerConfig) (*Process, error) {
	p := &Process{
		Version: cfg.Version,
		Session: cfg.Session,
	}
	return p, nil
}

func (p *Process) Active() (bool, error) {
	//TODO
	return false, nil
}

func (p *Process) ServiceName() string {
	return fmt.Sprintf("%s-%s", "mc-server", p.Version)
}

func (p *Process) Start() error {
	return p.Session.Run("systemctl", "start", p.ServiceName())
}

func (p *Process) Stop() error {
	return p.Session.Run("systemctl", "stop", p.ServiceName())
}

func (p *Process) Restart() error {
	return p.Session.Run("systemctl", "restart", p.ServiceName())
}
