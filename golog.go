package golog

import (
	"fmt"
	"os"
	"time"
)

const (
	RED   = "\x1B[31m" // red
	GRN   = "\x1B[32m" // green
	YEL   = "\x1B[33m" // yellow
	BLU   = "\x1B[34m" // blue
	MAG   = "\x1B[35m" // magenta
	CYN   = "\x1B[36m" // cyan
	WHT   = "\x1B[37m" // white
	RESET = "\x1B[0m"
)

var logWhite = false

func init() {
	if os.Getenv("LOG_WHITE") == "true" {
		logWhite = true
	}
}

func log(color, message string) {
	if logWhite {
		fmt.Println(message)
	} else {
		fmt.Println(color, message, RESET)
	}
}

// Info log
func Info(label, message string) {
	now := getDatetimeNow()
	log(BLU, fmt.Sprintf("[ %s | %s ] INFO - %s", now, label, message))
}

// I/O log
func Io(label, message string) {
	now := getDatetimeNow()
	log(MAG, fmt.Sprintf("[ %s | %s ] I/O - %s", now, label, message))
}

// Service log
func Service(label, message string) {
	now := getDatetimeNow()
	log(CYN, fmt.Sprintf("[ %s | %s ] SERVICE - %s", now, label, message))
}

// System log
func System(label, message string) {
	now := getDatetimeNow()
	log(GRN, fmt.Sprintf("[ %s | %s ] SYSTEM - %s", now, label, message))
}

// Error log
func Error(label string, message error) {
	now := getDatetimeNow()
	log(RED, fmt.Sprintf("[ %s | %s ] ERROR - %s", now, label, message.Error()))
}

// Warning log
func Warn(label, message string) {
	now := getDatetimeNow()
	log(YEL, fmt.Sprintf("[ %s | %s ] WARNING - %s", now, label, message))
}

func getDatetimeNow() string {
	now := time.Now().Local()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func NewLogTag(tag string) *_LogTag {
	return &_LogTag{
		tag: tag,
	}
}

type _LogTag struct {
	tag string
}

func (log *_LogTag) Info(method, message string) {
	Info(log.tag, fmt.Sprintf("(%s) %s", method, message))
}

func (log *_LogTag) Io(method, message string) {
	Io(log.tag, fmt.Sprintf("(%s) %s", method, message))
}

func (log *_LogTag) Service(method, message string) {
	Service(log.tag, fmt.Sprintf("(%s) %s", method, message))
}

func (log *_LogTag) System(method, message string) {
	System(log.tag, fmt.Sprintf("(%s) %s", method, message))
}

func (log *_LogTag) Error(method string, err error) {
	Error(log.tag, fmt.Errorf("(%s) %s", method, err.Error()))
}

func (log *_LogTag) Warn(method, message string) {
	Warn(log.tag, fmt.Sprintf("(%s) %s", method, message))
}
