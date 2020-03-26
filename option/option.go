package option

import "github.com/appnaconda/log"

func WithLevel(level log.Level) log.Option {
	return withLogLevel{level: level}
}

type withLogLevel struct {
	level log.Level
}

func (lv withLogLevel) Apply(logger log.Logger) {
	logger.SetLevel(lv.level)
}

func WithFormat(format log.Format) log.Option {
	return withLogFormat{format: format}
}

type withLogFormat struct {
	format log.Format
}

func (f withLogFormat) Apply(logger log.Logger) {
	logger.SetFormat(f.format)
}
