package manager

import (
	model2 "github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
	"reflect"
)

type Config struct {
	AllowListPath   string
	PermissionsPath string
}
type Manager struct {
	AllowListManager   *fm.FileManager[model2.AllowList]
	PermissionsManager *fm.FileManager[model2.Permissions]
}

func New(cfg *Config) *Manager {
	manager := &Manager{
		AllowListManager: fm.Default[model2.AllowList](cfg.AllowListPath),
		PermissionsManager: fm.New[model2.Permissions](&fm.Config{
			Path: cfg.PermissionsPath,
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
