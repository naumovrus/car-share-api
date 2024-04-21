package logger

import (
	"car-api/config"
	"fmt"
	"os"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger
type ApiLogger struct {
	cfg    *config.Config
	logger zerolog.Logger
}

// App Logger constructor
func NewApiLogger(cfg *config.Config) *ApiLogger {
	return &ApiLogger{cfg: cfg}
}

var loggerLevelMap = map[string]zerolog.Level{
	"debug":    zerolog.DebugLevel,
	"info":     zerolog.InfoLevel,
	"warn":     zerolog.WarnLevel,
	"error":    zerolog.ErrorLevel,
	"panic":    zerolog.PanicLevel,
	"fatal":    zerolog.FatalLevel,
	"noLevel":  zerolog.NoLevel,
	"disabled": zerolog.Disabled,
}

func (a *ApiLogger) InitLogger() error {
	var w zerolog.LevelWriter
	a.logger = log.With().Caller().Logger()
	if a.cfg.Logger.InFile {
		logFile, err := os.OpenFile(a.cfg.Logger.FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
		if err != nil {
			return err
		}
		w = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, logFile)
	} else {
		w = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	a.logger = zerolog.New(w).Level(loggerLevelMap[a.cfg.Logger.Level]).With().CallerWithSkipFrameCount(a.cfg.Logger.SkipFrameCount).Timestamp().Logger().Hook(a)
	return nil
}

func (a *ApiLogger) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if loggerLevelMap[a.cfg.Logger.Level] > level {
		return
	}
}

func (a *ApiLogger) Debug(msg string) {
	go a.logger.Debug().Msg(msg)
}

func (a *ApiLogger) Debugf(template string, args ...interface{}) {
	go a.logger.Debug().Msgf(template, args...)
}

func (a *ApiLogger) Info(msg string) {
	go a.logger.Info().Msg(msg)
}

func (a *ApiLogger) Infof(template string, args ...interface{}) {
	go a.logger.Info().Msgf(template, args...)
}

func (a *ApiLogger) Warn(msg string) {
	go a.logger.Warn().Msg(msg)
}

func (a *ApiLogger) Warnf(template string, args ...interface{}) {
	go a.logger.Warn().Msgf(template, args...)
}

func (a *ApiLogger) Error(err error) {
	go a.logger.Error().Msg(err.Error())
}

func (a *ApiLogger) Errorf(template string, args ...interface{}) {
	go a.logger.Error().Msgf(template, args...)
}

func (a *ApiLogger) Panic(msg string) {
	go a.logger.Panic().Msg(msg)
}

func (a *ApiLogger) Panicf(template string, args ...interface{}) {
	go a.logger.Panic().Msgf(template, args...)
}

func (a *ApiLogger) Fatal(msg string) {
	go a.logger.Fatal().Msg(msg)
}

func (a *ApiLogger) Fatalf(template string, args ...interface{}) {
	a.logger.Fatal().Msgf(template, args...)
}

func (a *ApiLogger) Tracef(s string, i ...interface{}) {
	go a.logger.Trace().Msgf(s, i...)
}

func (a *ApiLogger) ErrorFull(error error) {
	pc, _, line, _ := runtime.Caller(1)
	det := runtime.FuncForPC(pc)
	msg := fmt.Sprintf("ERROR:\n%s :: %d :: %s", det.Name(), line, error.Error())
	go a.logger.Error().Stack().Err(error).Msg(msg)
}
