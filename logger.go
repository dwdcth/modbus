package modbus

import (
	"fmt"
	"log"
	"os"
)

type LogSaver interface {
	Write(dir int, url string, station string, data []byte)
}

const (
	DirTx = iota + 1
	DirRx
)

type logger struct {
	prefix       string
	customLogger *log.Logger
}

func newLogger(prefix string, customLogger *log.Logger) (l *logger) {
	l = &logger{
		prefix:       prefix,
		customLogger: customLogger,
	}

	return
}

func (l *logger) Info(msg string) {
	l.write(fmt.Sprintf("%s [info]: %s\n", l.prefix, msg))

	return
}

func (l *logger) Infof(format string, msg ...interface{}) {
	l.write(fmt.Sprintf("%s [info]: %s\n", l.prefix, fmt.Sprintf(format, msg...)))

	return
}

func (l *logger) Warning(msg string) {
	l.write(fmt.Sprintf("%s [warn]: %s\n", l.prefix, msg))

	return
}

func (l *logger) Warningf(format string, msg ...interface{}) {
	l.write(fmt.Sprintf("%s [warn]: %s\n", l.prefix, fmt.Sprintf(format, msg...)))

	return
}

func (l *logger) Error(msg string) {
	l.write(fmt.Sprintf("%s [error]: %s\n", l.prefix, msg))

	return
}

func (l *logger) Errorf(format string, msg ...interface{}) {
	l.write(fmt.Sprintf("%s [error]: %s\n", l.prefix, fmt.Sprintf(format, msg...)))

	return
}

func (l *logger) Fatal(msg string) {
	l.Error(msg)
	os.Exit(1)

	return
}

func (l *logger) Fatalf(format string, msg ...interface{}) {
	l.Errorf(format, msg...)
	os.Exit(1)

	return
}

func (l *logger) write(msg string) {
	if l.customLogger == nil {
		os.Stdout.WriteString(msg)
	} else {
		l.customLogger.Print(msg)
	}

	return
}
