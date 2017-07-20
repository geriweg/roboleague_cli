package robocode

import (
	"os"
	"errors"
	"fmt"
)

type Config struct {
	baseDirectory string
}

/*var config = Config {
	baseDirectory: "s"
}
*/

func DirectoryExists(path string) error {
	fmt.Printf("func is called with: %s\n", path)
	return nil
}

func IsDir(season string) (error) {
	src, err := os.Stat(season)
	if err != nil {
		return err
	}

	if !src.IsDir() {
		return errors.New("Not a directory: " + season)
	}

	return nil
}