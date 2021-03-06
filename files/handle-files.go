package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// UppercaseDirs sets every directory under rootDir to uppercase
func UppercaseDirs(rootDir string) {
	loopDirs(&rootDir, func(i int, f os.FileInfo) error {
		newName := strings.ToUpper(f.Name())

		err := os.Rename(rootDir+f.Name(), rootDir+newName)
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
}

// LowercaseDirs sets every directory under rootDir to lowercase
func LowercaseDirs(rootDir string) {
	loopDirs(&rootDir, func(i int, f os.FileInfo) error {
		newName := strings.ToLower(f.Name())

		err := os.Rename(rootDir+f.Name(), rootDir+newName)
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
}

// EnumerateDirs renames directories inside rootDir making them start with numbers
// (eg. dirname = folder => 1-folder)
func EnumerateDirs(rootDir string, removeSpaces, upperCase, lowerCase bool) {
	var newName string
	loopDirs(&rootDir, func(i int, f os.FileInfo) error {
		if upperCase {
			newName = strings.ToUpper(f.Name())
		} else if lowerCase {
			newName = strings.ToLower(f.Name())
		} else {
			newName = f.Name()
		}

		if removeSpaces {
			newName = strings.ReplaceAll(newName, " ", "_")
		}

		if !IsDirEnumerated(newName) {
			newName = fmt.Sprintf("%-2d-%v", i, newName)
			newName = strings.ReplaceAll(newName, " ", "-")
		}

		prefix := newName[:2]
		orderedPrefix := fmt.Sprintf("%-2d", i)
		if prefix != orderedPrefix {
			newName = strings.Replace(newName, prefix, orderedPrefix, 1)
			newName = strings.ReplaceAll(newName, " ", "-")
		}

		err := os.Rename(rootDir+f.Name(), rootDir+newName)
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
}

// RemoveDirsEnumeration removes enumeration generated by the EnumerateDirs() function
func RemoveDirsEnumeration(rootDir string) {
	var newName string
	loopDirs(&rootDir, func(i int, f os.FileInfo) error {
		newName = f.Name()
		replaceRegex(`^[0-9]{1,2}-{1,2}`, &newName, "")

		err := os.Rename(rootDir+f.Name(), rootDir+newName)
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
}

// RemoveDirSubstrings removes substrings that match the regex pattern
func RemoveDirSubstrings(rootDir, pattern string) {

}

// IsDirEnumerated checks if a dirname already starts with 2 numbers
// (ex. 01-Dir = true, 0-Dir = false, Dir = false)
func IsDirEnumerated(dirname string) bool {
	matched, err := regexp.MatchString(`^[0-9]{1,2}`, dirname)

	if err != nil {
		log.Fatal(err)
	}

	return matched
}

type handleDir func(int, os.FileInfo) error

func loopDirs(rootDir *string, dirHandler handleDir) {
	validateRootDir(rootDir)
	files, err := ioutil.ReadDir(*rootDir)
	if err != nil {
		log.Fatal(err)
	}

	for i, f := range files {
		var err error
		if f.IsDir() {
			err = dirHandler(i, f)
		}

		if err != nil {
			continue
		}
	}
}

func validateRootDir(rootDir *string) {
	if (*rootDir)[len(*rootDir)-1] != '/' {
		(*rootDir) = (*rootDir) + "/"
	}
}

func replaceRegex(expression string, old *string, new string) {
	exp, err := regexp.Compile(expression)
	if err != nil {
		log.Fatal(err)
	}

	*old = exp.ReplaceAllString(*old, new)
}
