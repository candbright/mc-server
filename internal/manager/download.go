package manager

import (
	"fmt"
	"path"
)

type DownloadUtil struct {
	RootDir string
	Version string
}

func (m *DownloadUtil) DownloadUrl() string {
	return fmt.Sprintf("https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip", m.Version)
}

func (m *DownloadUtil) ZipFile() string {
	return fmt.Sprintf("bedrock-server-%s.zip", m.Version)
}

func (m *DownloadUtil) ZipExist() bool {
	//zip文件是否存在
	//TODO
	return true
}

func (m *DownloadUtil) ServerDir() string {
	return path.Join(m.RootDir, fmt.Sprintf("server-%s", m.Version))
}

func (m *DownloadUtil) ServerExist() bool {
	//服务器目录是否存在
	//TODO
	return true
}

func (m *DownloadUtil) Download() error {
	//1. 检测是否存在当前版本的服务器目录
	if m.ServerExist() {
		return nil
	}

	//2. 不存在当前版本的服务器目录，则检测是否存在当前版本的zip文件
	if !m.ZipExist() {
		//不存在zip文件，先下载文件
		//TODO
	}

	//3. 解压zip文件
	//TODO
	return nil
}

func (m *DownloadUtil) Delete() error {
	//1. 删除服务器目录
	if m.ServerExist() {
		//TODO
	}

	//2. 删除zip文件
	if m.ZipExist() {
		//TODO
	}
	return nil
}
