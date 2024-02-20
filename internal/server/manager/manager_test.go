package manager

import "testing"

func TestNew(t *testing.T) {
	manager := New(&Config{
		AllowListPath:   "./allowlist.json",
		PermissionsPath: "./server.properties",
	})
	t.Log(manager.PermissionsManager.Data.A)
	manager.PermissionsManager.Data.A = "test change"
	err := manager.PermissionsManager.Write()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(manager.PermissionsManager.Data.A)
}
