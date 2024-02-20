package manager

import (
	"fmt"
	"path"
)

type DownloadManager struct {
	RootDir string
	Version string
}

func (m *DownloadManager) DownloadUrl() string {
	return fmt.Sprintf("https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip", m.Version)
}

func (m *DownloadManager) ZipFile() string {
	return fmt.Sprintf("bedrock-server-%s.zip", m.Version)
}

func (m *DownloadManager) ZipExist() bool {
	//zip文件是否存在
	return true
}

func (m *DownloadManager) ServerDir() string {
	return path.Join(m.RootDir, fmt.Sprintf("server-%s", m.Version))
}

func (m *DownloadManager) ServerExist() bool {
	//服务器目录是否存在
	return true
}

func (m *DownloadManager) Download() error {
	//1. 检测是否存在当前版本的服务器目录
	if m.ServerExist() {
		return nil
	}

	//2. 不存在当前版本的服务器目录，则检测是否存在当前版本的zip文件
	if !m.ZipExist() {
		//不存在zip文件，先下载文件
	}

	//3. 解压zip文件
	return nil
}
