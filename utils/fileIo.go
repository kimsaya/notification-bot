package utils

import (
	"log"
	"os"
	"strings"
)

// CheckDirectoryOfFile check if Directory of file is available
func CheckDirectoryOfFile(filePath string) bool {
	fullPath := ""
	paths := strings.Split(filePath, "/")
	for i, path := range paths {
		if i == 0 {
			fullPath += path
			continue
		} else if i < len(paths)-1 {
			fullPath += "/" + path
		}
	}
	return CreateDirectory(fullPath)
}

// CreateDirectory create a directory
func CreateDirectory(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			log.Println("[ER] Directory Created:", err)
			return false
		}
	}
	return true
}

// OpenFile open a file for READ, WRITE, and APPEND
func OpenFile(filePath string) (refile *os.File) {
	CheckDirectoryOfFile(filePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		log.Println("[ER] Open File : ", err)
		return nil
	}
	return file
}
