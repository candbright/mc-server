package manager

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestManager_LatestVersion(t *testing.T) {
	manager := &Manager{
		client: resty.New(),
	}
	version, err := manager.LatestVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}
