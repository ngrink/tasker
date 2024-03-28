package lib

import (
	"fmt"
)

func GetDataBasepath() string {
	user := GetCurrentUser()

	switch user {
	case "root":
		return "/root/.tasker/"
	default:
		return fmt.Sprintf("/home/%s/.tasker/", user)
	}
}
