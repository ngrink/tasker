package lib

import (
	"log"
	"os/user"
)

func GetCurrentUser() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}
