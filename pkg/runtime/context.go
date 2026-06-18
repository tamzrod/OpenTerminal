package runtime

import (
	"os"
	"os/user"
)

type Context struct {
	WorkDir string
	User    string
}

func NewContext() Context {
	dir, _ := os.Getwd()

	currentUser, err := user.Current()
	username := ""

	if err == nil {
		username = currentUser.Username
	}

	return Context{
		WorkDir: dir,
		User:    username,
	}
}
