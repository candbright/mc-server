package fm

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	Path      string
	Marshal   func(v any) ([]byte, error)
	Unmarshal func(data []byte, v any) error
}

type FileManager[T any] struct {
	Cfg  *Config
	Data T
}

func Default[T any](path string) *FileManager[T] {
	cfg := &Config{
		Path:      path,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
	return New[T](cfg)
}

func New[T any](cfg *Config) *FileManager[T] {
	manager := &FileManager[T]{
		Cfg: cfg,
	}
	err := manager.Read()
	if err != nil {
		panic(err)
	}
	return manager
}

func (manager *FileManager[T]) Read() error {
	fileBytes, err := os.ReadFile(manager.Cfg.Path)
	if err != nil {
		return errors.WithStack(err)
	}
	err = manager.Cfg.Unmarshal(fileBytes, &manager.Data)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (manager *FileManager[T]) Write() error {
	marshalBytes, err := manager.Cfg.Marshal(manager.Data)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.WriteFile(manager.Cfg.Path, marshalBytes, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
