package logger

import (
	"log"
	"os"
	"strconv"
	"time"

	"umbrella/internal/channels"
)

const (
	LOG_FILE_NAME = "Umbrella"
)

var (
	CurrentLogFile *os.File
	FilePath       string
	TempDir        = os.Getenv("TEMP")
)

func Run() {
	FilePath = TempDir + "\\" + LOG_FILE_NAME + strconv.FormatInt(time.Now().Unix(), 10) + ".log"

	// Create a new log file if it doesn't exist
	_, err := os.Create(FilePath)
	if err != nil {
		log.Panicf("Failed to create log file: %v", err)
	}

	CurrentLogFile, err := os.OpenFile(FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Could not open log file: %v", err)
	}
	defer CurrentLogFile.Close()

	log.SetOutput(CurrentLogFile)

	log.Println("Log started at " + time.Now().String())

	for {
		msg := <-channels.Logs
		var timestamp = time.Now().UTC().Format("2006-01-02T15:04:05-0700")

		time.Sleep(1 * time.Millisecond)

		// Write the log message to the log file
		_, err := CurrentLogFile.WriteString(timestamp + " " + msg + "\r\n")
		if err != nil {
			log.Fatalf("Could not write to log file: %v", err)
		}

		time.Sleep(1 * time.Millisecond)
	}
}
