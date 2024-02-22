package model

type ServerInfo struct {
	Version  string `json:"version"`
	ZipExist bool   `json:"zip_exist"`
	DirExist bool   `json:"dir_exist"`
	Running  bool   `json:"running"`
}
