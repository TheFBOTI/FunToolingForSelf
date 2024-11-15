package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Make a wizard/GUI for this?
func main() {
	// Tell the application where to read from
	var inputDirectory string
	fmt.Println("Enter the path of the directory: ")
	_, err := fmt.Scan(&inputDirectory)
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

	// Tell the application if you want the files to be in a single top level directory or to copy the folder structure of the input directory
	var flatten bool
	fmt.Println("True or False for whether you want files in one directory or to have same directory structure as the input Directory : ")
	_, boolErr := fmt.Scan(&flatten)
	if boolErr != nil {
		fmt.Println(boolErr)
		return
	}

	// Iterate over the inputDirectory and find all isDir true and append them to an array with their absolute path
	directories := findDirectories(inputDirectory)

	// Grabs all files and folders in the directory.
	for _, directoryPath := range directories {
		arrayOfFiles, _ := os.ReadDir(directoryPath.Name)

		// Only used if flatten is false - but used to append to outPutDirectory Folder Creation and File Creation, extracts the difference between input and any folders found by findDirectories
		extension, _ := filepath.Rel(inputDirectory, directoryPath.Name)
		// Depending on flatten this will then check if a base Output Folder is present or if a nested Output Folder is present, if no it will create it with Read, Write, Execute permissions
		if len(extension) != 0 && !flatten {
			if _, err := os.Stat(outputDirectory + "/" + extension); os.IsNotExist(err) {
				os.MkdirAll(outputDirectory+"/"+extension, 0755)
			}
		} else {
			if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
				os.MkdirAll(outputDirectory, 0755)
			}
		}
		// Iterate over each file
		for _, file := range arrayOfFiles {
			//Check if it's a file or a folder - false is a file
			if !file.IsDir() {
				codeFile, err := os.Open(directoryPath.Name + "/" + file.Name())
				if err != nil {
					fmt.Println(err)
					return
				}
				defer codeFile.Close()

				// Extract the Content from original File(s)
				scanner := bufio.NewScanner(codeFile)
				var content string
				for scanner.Scan() {
					if scanner.Err() != io.EOF {
						content += scanner.Text() + "\n"
					} else {
						break
					}
				}

				// Create and write to File(s)
				var newFile *os.File
				var createFileError error

				// Depending on Flatten this will either write the contents from Scanner to either a file at base Output Directory or a nested view in line with the input directory
				if flatten {
					extensionStringModified := strings.ReplaceAll(extension, "\\", "_")
					newFile, createFileError = os.Create(outputDirectory + "/" + extensionStringModified + "_" + file.Name() + ".txt")
				} else {
					newFile, createFileError = os.Create(outputDirectory + "/" + extension + "/" + file.Name() + ".txt")
				}
				if createFileError != nil {
					fmt.Println(createFileError)
					return
				}
				defer newFile.Close()
				// Write the contents of the file to the new file
				newFile.WriteString(content)
			}
		}
	}
}

type Directory struct {
	Name string
}

func findDirectories(rootPath string) []Directory {
	var dirs []Directory

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		// Check if the directory is a nested directory and add it to the list
		dirs = append(dirs, Directory{Name: path})

		return nil
	})

	return dirs
}
