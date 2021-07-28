package appender

import (
	"go-git-contributions/file/reader"
	"io/ioutil"
	"strings"
)

// AddNewSliceElementsToFile given a slice of strings representing paths, stores them to the filesystem
func AddNewSliceElementsToFile(filePath string, newRepos []string)  {
	existingRepos := reader.ParseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringsSliceToFile(repos, filePath)
}

// joinSlices adds the element of the `new` slice  into the `existing` slice, only if not already there
func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

// sliceContains returns true if `slice` contains `value`
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// dumpStringsSliceToFile writes content to the file in path `filePath` (overwriting existing content)
func dumpStringsSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	err := ioutil.WriteFile(filePath, []byte(content), 0755)
	if err != nil {
		// error writing file
		panic(err)
	}
}
