package manager

import (
	"fmt"
	"path"
)

type VersionInfo struct {
	RootDir string
	Version string
}

func (m *VersionInfo) DownloadUrl() string {
	return fmt.Sprintf("https://minecraft.azureedge.net/bin-linux/bedrock-server-%s.zip", m.Version)
}

func (m *VersionInfo) ZipFile() string {
	return fmt.Sprintf("bedrock-server-%s.zip", m.Version)
}

func (m *VersionInfo) ZipFilePath() string {
	return path.Join(m.RootDir, m.ZipFile())
}

func (m *VersionInfo) ServerDir() string {
	return path.Join(m.RootDir, fmt.Sprintf("server-%s", m.Version))
}
