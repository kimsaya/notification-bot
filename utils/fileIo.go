package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CheckDirectoryOfFile check if Directory of file is available
func checkDirectoryOfFile(filePath string) bool {
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
	return createDirectory(fullPath)
}

// createDirectory create a directory
func createDirectory(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			log.Println("[ER] Directory Created:", err)
			return false
		}
	}
	return true
}

// removeFile remove a file
func removeFile(sourcePath string) bool {
	err := os.Remove(sourcePath)
	if err != nil {
		log.Println("[ER] RemoveFile:", err)
		return false
	}
	return true
}

// openFile open a file for READ, WRITE, and APPEND
func openFile(filePath string) (openedFile *os.File) {
	checkDirectoryOfFile(filePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		log.Println("[ER] OpenFile : ", err)
		return nil
	}
	return file
}

// ReadFile real all in file.
func ReadFile(filePath string) (result string, status bool) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("[ER] ioutil.ReadFile", err)
		return "", false
	}
	result = string(dat)
	return result, true
}

// WriteFile write to a file/ create if not exist
func WriteFile(filePath, message string) bool {
	configData := []byte(message)
	err := ioutil.WriteFile(filePath, configData, 0644)
	if err != nil {
		log.Println("[ER] ioutil.WriteFile:", err)
		return false
	}
	return true
}

// AppendFile append a file with string
func AppendFile(filePath, message string) bool {
	file, err := os.OpenFile(filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("[ER] os.OpenFile:", err)
		return false
	}
	defer file.Close()
	if _, err := file.WriteString(message); err != nil {
		log.Println("[ER] file.WriteString", err)
		return false
	}
	return true
}

// getFileNameFromPath splite out the full path then get the name
func getFileNameFromPath(path string) (fileName string) {
	sectors := strings.Split(path, "/")
	if len(sectors) == 1 {
		sectors = strings.Split(path, "\\")
	}
	return sectors[len(sectors)-1]
}

// getNoLimitInDirectory return list of all files in a directory
func getNoLimitInDirectory(directory string) (listFilePaths []string) {
	var paths []string
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println("[ER] GetNoLimitInDirectory:", err)
		createDirectory(directory)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		paths = append(paths, directory+"/"+f.Name())
	}
	return paths
}

// getLimitInDirectory return list of files in a directory with limit of items
func getLimitInDirectory(directory string, limit int) (listFilePaths []string) {
	var paths []string
	i := 0
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if i > 0 {
				paths = append(paths, path)
			}
			if i >= limit {
				return nil
			}
			i++
			return nil
		})
	if err != nil {
		log.Println("[ER] GetLimitInDirectory", err)
		createDirectory(directory)
	}
	return paths
}
