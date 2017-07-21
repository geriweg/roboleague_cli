package robocode

import (
	"os"
)

type fileSystem interface {
	Create(name string) (*os.File, error)
}

type osFS struct {}

func (osFS) Create(name string) (*os.File, error) {
	return os.Create(name)
}


