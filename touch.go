package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Exit status codes
const (
	successStatus = iota
	pathErrorStatus
	writeErrorStatus
	getWDFailErrorStatus
)

// The kind of operation to perform with the inputs
const (
	newFileOnly = iota + 1
	newFileAndContent
)

// The index of the positional argument in the flag parser
const (
	fileNamePosArgIndex = iota
	contentPosArgIndex
)

// Helpful string constants used in this program
const (
	usageMessage = `Usage: ./touch <name> [content]
	name: the name of the file to create
	content: initial content to put in the new file. 
	This command creates a new file. If content is provided it is placed in the 
	new file. If the path is absolute then the given path is unmodified and the 
	file is attempted to be created using that path. Otherwise the path is 
	prepended with the current working directory path and this new path is used 
	to create the file. If the file already exists its original content is 
	erased.`
	noContent           = ""
	successFormatString = "File \"%s\" was created without error.\n"
)

// printError: prints the given error if not nil and return true,
//				if the error is nil does nothing and returns false
func printError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

// getFileName: if file is an absolute path it is returned as is, otherwise
//				prepends the working directory to given path and returns it
func getFileName() string {
	fileName := string(flag.Arg(fileNamePosArgIndex))
	if filepath.IsAbs(fileName) {
		return fileName
	}
	wd, err := os.Getwd()
	if printError(err) {
		os.Exit(getWDFailErrorStatus)
	}
	return filepath.Join(wd, fileName)
}

// success: prints success statement to stdout and exits with status code of 0
func success(fileName string) {
	fmt.Printf(successFormatString, fileName)
	os.Exit(successStatus)
}

// printUsageMessage: prints the usage message for the tool.
func printUsageMessage() {
	fmt.Println(usageMessage)
}

// createNewFile: creates a new file with the arguments provided in flag,
// 					if there is content it is added to the file before closing.
//					The file is properly closed and an exit status is reported
//					in all cases.
func createNewFile(content string) {
	fileName := getFileName()
	file, err := os.Create(fileName)
	if printError(err) {
		file.Close()
		os.Exit(pathErrorStatus)
	}
	if content != noContent {
		_, err = file.WriteString(content)
		if printError(err) {
			file.Close()
			os.Exit(writeErrorStatus)
		}
	}
	file.Close()
	success(fileName)
}

// main: the entry point for the tool, determines which action to take based on
//			the given arguments and prints a usage message if invalid arguments
//			are provided.
func main() {
	flag.Usage = printUsageMessage
	flag.Parse()
	switch flag.NArg() {
	case newFileOnly:
		createNewFile(noContent)
	case newFileAndContent:
		createNewFile(flag.Arg(contentPosArgIndex))
	default:
		flag.Usage()
		os.Exit(successStatus)
	}
}
