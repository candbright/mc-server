package manager

import (
	"fmt"
	"path"

	"github.com/candbright/go-ssh/ssh"
)

type DownloadUtil struct {
	RootDir string
	Version string
	Session ssh.Session
}

func (m *DownloadUtil) DownloadUrl() string {
	return fmt.Sprintf("https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip", m.Version)
}

func (m *DownloadUtil) ZipFile() string {
	return fmt.Sprintf("bedrock-server-%s.zip", m.Version)
}

func (m *DownloadUtil) ZipFilePath() string {
	return path.Join(m.RootDir, m.ZipFile())
}

func (m *DownloadUtil) ZipExist() (bool, error) {
	//zip文件是否存在
	return m.Session.Exists(m.ZipFilePath())
}

func (m *DownloadUtil) ServerDir() string {
	return path.Join(m.RootDir, fmt.Sprintf("server-%s", m.Version))
}

func (m *DownloadUtil) ServerExist() (bool, error) {
	//服务器目录是否存在
	return m.Session.Exists(m.ServerDir())
}

func (m *DownloadUtil) Download() error {
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

func (m *DownloadUtil) Delete() error {
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
