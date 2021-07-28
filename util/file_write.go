package util

import (
	"bufio"
	"io/ioutil"
	"os"
)

func WriteText(filePath, content string, append bool) error {
	var file *os.File
	var err error
	if append {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	} else {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	}
	if err != nil {
		return err
	}
	defer file.Close()
	outputWriter := bufio.NewWriter(file)
	outputWriter.WriteString(content)
	outputWriter.Flush()
	return nil
}


func ReadText(inputFilePath string) (string, error) {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return "", err
	}
	sql, err := ioutil.ReadAll(inputFile)
	
	return string(sql), nil
}