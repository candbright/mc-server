package core

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/server-mc/internal/repo"
	"github.com/candbright/server-mc/pkg/fm"
	"github.com/pkg/errors"
	"html/template"
	"path"
)

//go:embed template
var tmpl embed.FS

const (
	serviceDir     = "/opt/bin"
	environmentDir = "/etc/sysconfig"
	allowListFile  = "allowlist.json"
	propertiesFile = "server.properties"
)

type ServerConfig struct {
	Version string
	RootDir string
	Session ssh.Session
}

type Server struct {
	Version              string
	RootDir              string
	Session              ssh.Session
	Process              *Process
	AllowListFile        *fm.FileManager[repo.AllowList]
	ServerPropertiesFile *fm.FileManager[map[string]string]
}

func NewServer(cfg *ServerConfig) (*Server, error) {
	process, err := NewProcess(cfg)
	if err != nil {
		return nil, err
	}
	server := &Server{
		Version: cfg.Version,
		RootDir: cfg.RootDir,
		Session: cfg.Session,
		Process: process,
	}
	exist, err := server.ServerExist()
	if err != nil {
		return nil, err
	}
	if !exist {
		return server, nil
	}
	err = server.Reload()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (server *Server) DownloadUrl() string {
	return fmt.Sprintf("https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip", server.Version)
}

func (server *Server) ZipFileName() string {
	return fmt.Sprintf("bedrock-server-%s.zip", server.Version)
}

func (server *Server) ZipFile() string {
	return path.Join(server.RootDir, server.ZipFileName())
}

func (server *Server) ZipExist() (bool, error) {
	//zip文件是否存在
	return server.Session.Exists(server.ZipFile())
}

func (server *Server) ServerDirName() string {
	return fmt.Sprintf("server-%s", server.Version)
}

func (server *Server) ServerDir() string {
	return path.Join(server.RootDir, server.ServerDirName())
}

func (server *Server) ServerExist() (bool, error) {
	//服务器目录是否存在
	return server.Session.Exists(server.ServerDir())
}

func (server *Server) ExecFile() string {
	return path.Join(server.ServerDir(), "bedrock_server")
}

func (server *Server) Reload() error {
	server.ServerPropertiesFile = ServerPropertiesFM(server.Version, server.RootDir)
	server.AllowListFile = fm.Default[repo.AllowList](path.Join(server.RootDir, allowListFile))
	err := server.PrepareService()
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Download() error {
	//检测是否存在当前版本的服务器目录
	existS, err := server.ServerExist()
	if err != nil {
		return err
	}
	if existS {
		return nil
	}
	//不存在当前版本的服务器目录，则检测是否存在当前版本的zip文件
	existZ, err := server.ZipExist()
	if err != nil {
		return err
	}
	//不存在当前版本的zip文件，先下载
	if !existZ {
		err = server.Session.Run("wget", server.DownloadUrl(), "-P", server.RootDir)
		if err != nil {
			return err
		}
	}

	//解压zip文件
	err = server.Session.MakeDirAll(server.ServerDir(), 0777)
	if err != nil {
		return err
	}
	err = server.Session.Run("unzip", server.ZipFile(), "-d", server.ServerDir())
	if err != nil {
		return err
	}
	//4. reload
	err = server.Reload()
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Delete() error {
	//如果服务器正在运行则先关闭
	active, err := server.Process.Active()
	if err != nil {
		return err
	}
	if active {
		err = server.Process.Stop()
		if err != nil {
			return err
		}
	}
	//删除服务器目录
	//TODO: 备份
	existS, err := server.ServerExist()
	if err != nil {
		return err
	}
	if existS {
		err = server.Session.RemoveAll(server.ServerDir())
		if err != nil {
			return err
		}
	}
	//删除zip文件
	existZ, err := server.ZipExist()
	if err != nil {
		return err
	}
	if existZ {
		err = server.Session.Remove(server.ZipFile())
		if err != nil {
			return err
		}
	}
	//删除service
	err = server.Session.Remove(server.Process.ServiceName())
	if err != nil {
		//TODO
	}
	err = server.Session.Remove(path.Join(environmentDir, server.Process.ServiceName()))
	if err != nil {
		//TODO
	}
	return nil
}

func (server *Server) PrepareService() error {
	content, err := template.ParseFS(tmpl,
		path.Join("template", "service"))
	if err != nil {
		return errors.WithStack(err)
	}
	var result bytes.Buffer
	err = content.Execute(&result, struct {
		ServiceName     string
		EnvironmentFile string
		ExecStart       string
	}{
		ServiceName:     server.Process.ServiceName(),
		EnvironmentFile: path.Join(environmentDir, server.Process.ServiceName()),
		ExecStart:       server.ExecFile(),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	err = server.Session.WriteString(path.Join(serviceDir, server.Process.ServiceName()), result.String())
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) AddUser(username string) error {
	//TODO
	return nil
}

func (server *Server) DeleteUser(id string) error {
	//TODO
	return nil
}
