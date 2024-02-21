package process

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"
	"path"

	"github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
)

//go:embed template
var serverPropertiesTmpl embed.FS

const (
	PropertiesFile = "server.properties"
)

type ServerProperties struct {
	FileManager *fm.FileManager[model.Properties_1_20_62_02]
}

func NewServerProperties(dir string) *ServerProperties {
	fileManager := fm.New[model.Properties_1_20_62_02](&fm.Config{
		Path:    path.Join(dir, PropertiesFile),
		Marshal: Encode,
		Unmarshal: func(data []byte, v any) error {
			p := properties.MustLoadString(string(data))
			err := p.Decode(v)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		},
	})
	return &ServerProperties{
		FileManager: fileManager,
	}
}

func Encode(v any) ([]byte, error) {
	content, err := template.ParseFS(serverPropertiesTmpl, path.Join("template", PropertiesFile))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result bytes.Buffer
	err = content.Execute(&result, v)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result.Bytes(), nil
}
