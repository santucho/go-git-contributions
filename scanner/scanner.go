package scanner

import (
	"fmt"
	"go-git-contributions/dotfile"
	"go-git-contributions/file/appender"
	"log"
	"os"
	"strings"
)

// Scan given a path crawls it and its subfolders
// searching for Git repositories
func Scan(folder string) {
	fmt.Printf("found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filePath := dotfile.GetDotFilePath()
	appender.AddNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

// recursiveScanFolder starts the recursive search of git repositories
// living in the `folder` subtree
func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func scanGitFolders(folders []string, folder string) []string {

	// trip the last `/`
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatalln(err)
	}
	files, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}

			// It explicitly avoids going into folders called vendor or node_modules since those folders can be huge
			// and usually you don’t put your Git repositories in there, we can safely ignore them.
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}
