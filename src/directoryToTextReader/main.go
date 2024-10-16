package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TODO: Make a wizard/GUI for this rather
func main() {
	// Tell the application where to read from
	var dirPath string
	fmt.Println("Enter the path of the directory: ")
	_, err := fmt.Scan(&dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Tell the application where you want to put the output file(s)
	var outputDirectory string
	fmt.Println("Enter the path of the output directory: ")
	_, inputError := fmt.Scan(&outputDirectory)
	if inputError != nil {
		fmt.Println(inputError)
		return
	}

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
					content += scanner.Text() + "\n"
				} else {
					break
				}
			}

			newFile, err := os.Create(outputDirectory + file.Name() + ".txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer newFile.Close()
			// Write the contents of the file to the new file
			newFile.WriteString(content)
		}
	}
}
