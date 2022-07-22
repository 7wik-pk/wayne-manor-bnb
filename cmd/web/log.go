package web

import (
	"io"
	"os"
)

var file *os.File

func initLogFile(filepath string) (err error) {
	file, err = os.Create(filepath)
	if err != nil {
		return err
	}

	return nil
}

func getLogWriter(shouldWriteToFile bool, shouldPrintToConsole bool) (logWriter io.Writer, err error) {

	var writer []io.Writer

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
