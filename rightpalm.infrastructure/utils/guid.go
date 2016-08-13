package utils

import (
	"log"
	"os/exec"
)

func NewGuid() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
