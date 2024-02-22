package core

import (
	"fmt"
	"testing"
)

func TestNewServerProperties(t *testing.T) {
	servers := New(&Config{
		LatestVersion: "1.20.62.02",
		RootDir:       "./example",
	})
	t.Log(servers.Latest.ServerPropertiesFile.Data["server-name"])
}

func TestServerPropertiesWrite(t *testing.T) {
	servers := New(&Config{
		LatestVersion: "1.20.62.02",
		RootDir:       "./example",
	})
	t.Log(servers.Latest.ServerPropertiesFile.Data["server-name"])
	servers.Latest.ServerPropertiesFile.Data["server-name"] = "NewServer"
	err := servers.Latest.ServerPropertiesFile.Write()
	if err != nil {
		t.Fatal(fmt.Sprintf("%+v", err))
	}
	t.Log(servers.Latest.ServerPropertiesFile.Data["server-name"])
}
