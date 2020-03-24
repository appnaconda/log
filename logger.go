package logger

import (
	"io"
	"strings"
)

type Option interface {
	Apply(logger Logger) error
}

type Fields map[string]interface{}

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

func (s Level) String() string {
	return levelName[s]
}

var levelName = []string{
	"debug",
	"info",
	"warn",
	"error",
}

var levelValue = map[string]Level{
	"debug": Debug,
	"info":  Info,
	"warn":  Warn,
	"error": Error,
}

func ParseLevel(l string) (Level, bool) {
	v, ok := levelValue[strings.ToLower(l)]
	return v, ok
}

type Format int

func (f Format) String() string {
	return formatName[f]
}

const (
	Json Format = iota
	Text
)

var formatName = [...]string{
	"json",
	"text",
}

var formatValue = map[string]Format{
	"json": Json,
	"text": Text,
}

func ParseFormat(f string) (Format, bool) {
	v, ok := formatValue[strings.ToLower(f)]
	return v, ok
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	With(fields Fields) Logger

	SetLevel(level Level) error
	SetOutput(output io.Writer)
	SetOutputFormat(format Format) error
}
