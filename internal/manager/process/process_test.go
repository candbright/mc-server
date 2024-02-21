package process

import (
	"fmt"
	"testing"
)

func TestNewServerProperties(t *testing.T) {
	process := New(&Config{
		RootDir: ".",
	})
	t.Log(process.ServerProperties.FileManager.Data.ServerName)
}

func TestServerPropertiesWrite(t *testing.T) {
	process := New(&Config{
		RootDir: ".",
	})
	process.ServerProperties.FileManager.Data.ServerName = "NewServer"
	err := process.ServerProperties.FileManager.Write()
	if err != nil {
		t.Fatal(fmt.Sprintf("%+v", err))
	}
	t.Log(process.ServerProperties.FileManager.Data.ServerName)
}
