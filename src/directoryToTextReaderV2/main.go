package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
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

	// Ensure the output directory exists
	if err := os.MkdirAll(outputDirectory, 0755); err != nil {
		panic(err)
	}

	// Traverse the input directory recursively
	if err := filepath.Walk(inputDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Determine the relative path from the input directory
		relPath, err := filepath.Rel(inputDirectory, path)
		if err != nil {
			return err
		}

		// Build the output file path
		outputPath := filepath.Join(outputDirectory, relPath)

		// Create the necessary directory structure in the output
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		// Open the source file
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// Create the destination file
		dstFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		// Copy content from source to destination
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		fmt.Printf("Copied: %s → %s\n", path, outputPath)
		return nil
	}); err != nil {
		panic(err)
	}

	fmt.Println("All files copied successfully.")
}
