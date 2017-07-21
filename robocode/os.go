package robocode

import (
	_ "io"
	"os"
	_ "fmt"
	"fmt"
)

type fileSystem interface {
	Create(name string) (*os.File, error)
}

type osFS struct {}

func (osFS) Create(name string) (*os.File, error) {
	fmt.Println("real Create")
	return os.Create(name)
}

