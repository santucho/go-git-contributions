package dotfile

import (
	"log"
	"os/user"
)

// getDotFilePath returns the dot file for the repos list.
// Creates it and the enclosing folder if it does not exist.
func GetDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

