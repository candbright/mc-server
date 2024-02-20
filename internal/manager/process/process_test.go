package process

import "testing"

func TestNew(t *testing.T) {
	process := New(&Config{
		RootDir: ".",
	})
	t.Log(process.ServerProperties.FileManager.Data.A)
	process.ServerProperties.FileManager.Data.A = "test change"
	err := process.ServerProperties.FileManager.Write()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(process.ServerProperties.FileManager.Data.A)
}
