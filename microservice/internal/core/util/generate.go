package util

import (
	"log"
	"os/exec"
)

func GenerateUUIDUnique() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(newUUID)
}
