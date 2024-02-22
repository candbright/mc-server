package core

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"path"

	"github.com/candbright/server-mc/pkg/fm"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
)

func ServerPropertiesFM(version, dir string) *fm.FileManager[map[string]string] {
	return fm.New[map[string]string](&fm.Config{
		Path: path.Join(dir, propertiesFile),
		Marshal: func(v any) ([]byte, error) {
			content, err := template.ParseFS(tmpl,
				path.Join("template", fmt.Sprintf("%s-%s", version, propertiesFile)))
			if err != nil {
				return nil, errors.WithStack(err)
			}
			var result bytes.Buffer
			err = content.Execute(&result, v)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return result.Bytes(), nil
		},
		Unmarshal: func(data []byte, v any) error {
			var err error
			p := properties.MustLoadString(string(data))
			if mapPtr, ok := v.(*map[string]string); ok {
				*mapPtr = p.Map()
				return nil
			}
			err = p.Decode(v)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		},
	})
}
