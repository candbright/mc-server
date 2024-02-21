package process

import (
	"bytes"
	"embed"
	_ "embed"
	model2 "github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
	"html/template"
	"path"
)

//go:embed template
var serverPropertiesTmpl embed.FS

const (
	PropertiesFile = "server.properties"
)

type ServerProperties struct {
	FileManager *fm.FileManager[model2.Properties]
}

func NewServerProperties(dir string) *ServerProperties {
	fileManager := fm.New[model2.Properties](&fm.Config{
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
