package gokit

import (
	"io"
	"io/fs"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Yaml[T any] struct{}

func (x *Yaml[T]) Encode(data any) (value []byte, err error) {
	return yaml.Marshal(data)
}

func (x *Yaml[T]) Decode(b []byte) (value T, err error) {
	vType := reflect.TypeOf(value)

	switch vType.Kind() {
	case reflect.Pointer:
		value = reflect.New(vType.Elem()).Interface().(T)
		err = yaml.Unmarshal(b, value)
	default:
		err = yaml.Unmarshal(b, &value)
	}

	return
}

func (x *Yaml[T]) Read(reader io.Reader) (value T, err error) {
	var b []byte

	if b, err = io.ReadAll(reader); err != nil {
		return
	}

	value, err = x.Decode(b)

	return
}

func (x *Yaml[T]) ReadFile(name string) (value T, err error) {
	var file *os.File

	if file, err = os.Open(name); err != nil {
		return
	}

	defer file.Close()

	value, err = x.Read(file)

	return
}

func (x *Yaml[T]) ReadFileFS(name string, fsys fs.FS) (value T, err error) {
	var file fs.File

	if file, err = fsys.Open(name); err != nil {
		return
	}

	defer file.Close()

	value, err = x.Read(file)

	return
}

func (x *Yaml[T]) GetConfigFile(paths ...any) (value T, err error) {
	var name String

	if name, err = Path.FromExecutable(paths...); err != nil {
		return
	}

	value, err = x.ReadFile(string(name))

	return
}
