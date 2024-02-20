package process

import (
	model2 "github.com/candbright/server-mc/internal/model"
	"github.com/candbright/server-mc/pkg/fm"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
	"path"
	"reflect"
)

const (
	PropertiesFile = "server.properties"
)

type ServerProperties struct {
	FileManager *fm.FileManager[model2.Properties]
}

func NewServerProperties(dir string) *ServerProperties {
	fileManager := fm.New[model2.Properties](&fm.Config{
		Path: path.Join(dir, PropertiesFile),
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
	})
	return &ServerProperties{
		FileManager: fileManager,
	}
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
