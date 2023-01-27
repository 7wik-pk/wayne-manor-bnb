package web

import (
	"errors"
	"io"
	"log"
	"os"
)

var file *os.File

func initLogFile(filepath string) (err error) {
	file, err = os.Create(filepath)
	if err != nil {
		log.Println("os.Create() err: ", err.Error())
		return err
	}

	return nil
}

func getLogWriter(shouldWriteToFile bool, shouldPrintToConsole bool) (logWriter io.Writer, err error) {

	var writer []io.Writer

	if shouldWriteToFile && file == nil {
		return nil, errors.New("file pointer is nil but shouldWriteToFile is true")
	}

	if shouldWriteToFile {
		writer = append(writer, file)
	}

	if shouldPrintToConsole {
		writer = append(writer, os.Stdout)
	}

	return io.MultiWriter(writer...), nil

}

func closeLogFile() {
	file.Close()
}
