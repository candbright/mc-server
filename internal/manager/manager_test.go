package manager

import (
	"testing"
)

var testManager = New(&Config{
	Version: "1.20.62.02",
	RootDir: "/home/lighthouse",
	/*SSHOptions: []options.Option{
		options.RemoteHostPWD("", 22, "", ""),
	},*/
})

func TestManager_LatestVersion(t *testing.T) {
	version, err := testManager.LatestVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}

func TestDownloadUtil_ZipExist(t *testing.T) {
	exist, err := testManager.ZipExist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(exist)
}

func TestManager_Download(t *testing.T) {
	err := testManager.Download()
	if err != nil {
		t.Fatal(err)
	}
}

func TestManager_Upgrade(t *testing.T) {
	var oldManager = New(&Config{
		Version: "1.20.15.01",
		RootDir: "/home/lighthouse",
		/*SSHOptions: []options.Option{
			options.RemoteHostPWD("", 22, "", ""),
		},*/
	})
	err := oldManager.Upgrade()
	if err != nil {
		t.Fatal(err)
	}
}
