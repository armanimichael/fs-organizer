package cli

import (
	"flag"
	"fmt"
	"fsorganizer/files"
	"os"
)

// InitCLI initializes CLI Commands
func InitCLI() {
	rootDir := flag.String("root", "./", "Specifies the root directory, current otherwise.")
	enumDirs := flag.Bool("enum", false, "Rename directories under rootDir to start with a number.")
	uppercase := flag.Bool("uppercase", false, "Converts every directories' names to UPPERCASE.")
	lowercase := flag.Bool("lowercase", false, "Converts every directories' names to lowercase.")
	noSpaces := flag.Bool("nospaces", false, "Removes white spaces from directories' names.")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.PrintDefaults()
	}

	if *uppercase && *lowercase {
		fmt.Println("Directories' names can't be both uppercase and lowercase.")
		os.Exit(1)
	}

	if *enumDirs {
		files.EnumerateDirs(*rootDir, *noSpaces, *uppercase, *lowercase)
	} else if *uppercase {
		files.UppercaseDirs(*rootDir)
	} else if *lowercase {
		files.LowercaseDirs(*rootDir)
	}
}
