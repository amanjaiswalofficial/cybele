package utils

import (
	"io/ioutil"
	"os"
	"log"
	"strings"
)

// ReadFileFromPath takes input as a path to read file from
// returns: error if exists, otherwise returns file
func ReadFileFromPath(path string) ([]byte, error) {
	readData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return readData, nil
}

// HandleError is used to ensure single place to handle all errors
// After displaying the message, exits the program
func HandleError(err string) {
	LogMessage(err)
	os.Exit(1)
}

// LogMessage is used to log messages to the console
func LogMessage(messages ...string) {
	if os.Getenv("LOGGING") == "true" {
		log.Println(strings.Join(messages, " "))
	}
}