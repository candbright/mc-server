package process

import (
	model2 "github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
	"path"
	"reflect"
)

type Config struct {
	RootDir        string
	AllowListFile  string
	PropertiesFile string
}

type Process struct {
	AllowListFile  *fm.FileManager[model2.AllowList]
	PropertiesFile *fm.FileManager[model2.Properties]
}

func New(cfg *Config) *Process {
	manager := &Process{
		AllowListFile: fm.Default[model2.AllowList](path.Join(cfg.RootDir, cfg.AllowListFile)),
		PropertiesFile: fm.New[model2.Properties](&fm.Config{
			Path: path.Join(cfg.RootDir, cfg.PropertiesFile),
			Marshal: func(v any) ([]byte, error) {
				newProperties := Encode(v)
				return []byte(newProperties.String()), nil
			},
			Unmarshal: func(data []byte, v any) error {
				p := properties.MustLoadString(string(data))
				err := p.Decode(v)
				if err != nil {
					return errors.WithStack(err)
				}
				return nil
			},
		}),
	}
	return manager
}

func Encode(v any) *properties.Properties {
	newProperties := properties.NewProperties()
	structType := reflect.TypeOf(v)
	structValue := reflect.ValueOf(v)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i).Interface()
		tag := field.Tag.Get("properties")
		err := newProperties.SetValue(tag, value)
		if err != nil {
			return newProperties
		}
	}
	return newProperties
}
