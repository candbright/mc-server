package process

import "testing"

func TestNew(t *testing.T) {
	process := New(&Config{
		RootDir:        ".",
		AllowListFile:  "allowlist.json",
		PropertiesFile: "server.properties",
	})
	t.Log(process.PropertiesFile.Data.A)
	process.PropertiesFile.Data.A = "test change"
	err := process.PropertiesFile.Write()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(process.PropertiesFile.Data.A)
}
