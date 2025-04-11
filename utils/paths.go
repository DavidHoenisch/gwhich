package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func GetUserConfigPath() (string, error) {

	home, err := os.UserConfigDir()
	log.Println(home)

	if err != nil {
		log.Println(err)
	}

	XDG_CONFIG_PATH := fmt.Sprintf("%s/%s", home, "gwhich")
	log.Println(XDG_CONFIG_PATH)

	_, err = os.Stat(XDG_CONFIG_PATH)

	if errors.Is(err, os.ErrNotExist) {
		return "", os.ErrNotExist
	}

	return XDG_CONFIG_PATH, nil
}
