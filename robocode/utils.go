package robocode

import "os/user"

func GetBaseDirectory() (string, error) {
	_, err := user.Current()
	if err != nil {
		return nil, err
	}

	return "baseDirectory", nil
}
