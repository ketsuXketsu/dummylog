package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// dummylog base struct
type Logger struct {
	outputDir string
}

func LogBase(l *Logger) (*os.File, error) {
	logFilePath := filepath.Join(l.outputDir, "logs")
	err := os.MkdirAll(logFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	logFileName := "log"
	logFilePath = filepath.Join(logFilePath, logFileName)

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	return f, nil
}

// Writes a message to a file, preppended by the current date and time
func (l *Logger) LogT(a ...any) {
	f, err := LogBase(l)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Print(a...)
}

// Writes a message to a file
func (l *Logger) Log(a ...any) {
	f, err := LogBase(l)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer f.Close()

	fmt.Fprint(f, a...)
}

/*

	-- Demo --

	dir, err := os.Getwd()
	if err != nil {
		return
	}

	logger := Logger{outputDir: dir}
	logger.LogT("Hello, " + "World!")
	logger.Log("Hello, " + "World!")


*/
