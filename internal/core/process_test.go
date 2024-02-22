package core

import (
	"github.com/candbright/go-ssh/ssh"
	"testing"
)

func TestProcess_Active(t *testing.T) {
	session, err := ssh.NewSession()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	process := &Process{
		Session: session,
	}
	active, err := process.Active()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log(active)
}

func TestProcess_Start(t *testing.T) {
	session, err := ssh.NewSession()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	process := &Process{
		Session: session,
	}
	err = process.Start()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	active, err := process.Active()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log(active)
}
