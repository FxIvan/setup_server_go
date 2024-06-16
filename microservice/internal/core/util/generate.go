package util

import (
	"log"
	"os/exec"
	"strings"
)

func GenerateUUIDUnique() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(newUUID))
}
