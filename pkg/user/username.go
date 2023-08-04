package user

import "os/user"

func GetUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		return "unknown_username"
	}

	return currentUser.Username
}
