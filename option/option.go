package option

import "github.com/appnaconda/logger"

func WithLevel(level logger.Level) logger.Option {
	return withLogLevel{level: level}
}

type withLogLevel struct {
	level logger.Level
}

func (lv withLogLevel) Apply(logger logger.Logger) {
	logger.SetLevel(lv.level)
}

func WithFormat(format logger.Format) logger.Option {
	return withLogFormat{format: format}
}

type withLogFormat struct {
	format logger.Format
}

func (f withLogFormat) Apply(logger logger.Logger) {
	logger.SetOutputFormat(f.format)
}
