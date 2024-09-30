package file

import (
	"fmt"
	"os"
	"regexp"
)

// ReadArtFile reads all the lines in the given ascii art file, then returns all the lines read.
func ReadArtFile(asciiArtFile string) []string {
	// Read the ascii art file.
	asciiArt, err := os.ReadFile(asciiArtFile)
	if err != nil {
		fmt.Printf("Error reading file: %q\n%v\n", asciiArtFile, err)
		os.Exit(1)
	}

	if len(asciiArt) == 0 {
		fmt.Println("Empty banner file")
		os.Exit(1)
	}

	fileInfo, err := os.Stat(asciiArtFile)
	if err != nil {
		fmt.Printf("Error getting file info: %q\n%v\n", asciiArtFile, err)
		os.Exit(1)
	}

	fileSize := fileInfo.Size()

	switch asciiArtFile {
	case "standard.txt":
		if fileSize != 6623 {
			fmt.Println("Standard Filesize Tampered")
			os.Exit(1)
		}
	case "shadow.txt":
		if fileSize != 7463 {
			fmt.Println("Shadow filesize Tampered")
			os.Exit(1)
		}
	case "thinkertoy.txt":
		if fileSize != 5558 {
			fmt.Println("Thinkertoy filesize Tampered")
			os.Exit(1)
		}
	case "lean.txt":
		if fileSize != 10871 {
			fmt.Println("lean filesize Tampered")
			os.Exit(1)
		}
	}

	// Convert asciiArt from bytes to string.
	asciiArtString := string(asciiArt)

	// Split asciiArtString into lines.
	re := regexp.MustCompile(`\r?\n`)
	bannerFileContents := re.Split(asciiArtString, -1)

	return bannerFileContents
}
