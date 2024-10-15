package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TODO: Make a wizard/GUI for this rather
func main() {
	// Iterate over all the files in the directory
	//TODO: turn directoryPath into a feedable path.
	dirPath := "DummyDirectory:/ToFolder"

	// Grabs all files and folders in the directory.
	arrayOfFiles, _ := os.ReadDir(dirPath)

	// Iterate over each file
	for _, file := range arrayOfFiles {
		//Tell the user what file is being checked
		fmt.Println(file.Name())
		//Check if it's a file or a folder
		if !file.IsDir() {
			// Enter if it's a file in the directory given
			codeFile, err := os.Open(dirPath + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			defer codeFile.Close()
			scanner := bufio.NewScanner(codeFile)
			var content string
			for scanner.Scan() {
				if scanner.Err() != io.EOF {
					content += scanner.Text()
				} else {
					break
				}
			}

		}
	}
}