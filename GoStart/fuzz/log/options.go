package log

import (
	"fmt"
	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	FORAMT_CONSOLE = "console"
	FORAMT_JSON    = "json"
	OUTPUT_STD     = "stdout"
	OUTPUT_STD_ERR = "stderr"

	flagLevel = "log.level"
)

type Options struct {
	OutputPaths     []string `json:"output-paths" mapstructure:"output-paths"`
	ErrorOuputPaths []string `json:"error-output-paths" mapstructure:"error-output-paths"`
	Level           string   `json:"level" mapstructure:"level"`
	Format          string   `json:"format" mapstructure:"format"`
	Name            string   `json:"name" mapstructure:"name"`
}

type Option func(o *Options)

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Level:           zapcore.InfoLevel.String(),
		Format:          FORAMT_CONSOLE,
		OutputPaths:     []string{OUTPUT_STD},
		ErrorOuputPaths: []string{OUTPUT_STD_ERR},
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

func WithLevel(level string) Option {
	return func(o *Options) {
		o.Level = level
	}
}

//就可以自定去定义检测规则
func (o *Options) Validate() []error {
	var errs []error
	format := strings.ToLower(o.Format)
	if format != FORAMT_CONSOLE && format != FORAMT_JSON {
		errs = append(errs, fmt.Errorf("not suppor format %s", o.Format))
	}
	return errs
}

//可以自己将options具体的列映射到flag的字段上
func (o *Options) AddFlags(fs pflag.FlagSet) {
	fs.StringVar(&o.Level, flagLevel, o.Level, "log level")
}
