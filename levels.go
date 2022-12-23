package logger

import "errors"

var levelError = errors.New("invalid level")

type Level int

const (
	Trace        Level = 0
	Debug        Level = 1
	Info         Level = 2
	Warn         Level = 3
	Error        Level = 4
	Fatal        Level = 5
	InvalidLevel Level = -1
)

var levelNames = map[Level]string{
	Trace: "trace",
	Debug: "debug",
	Info:  "info",
	Warn:  "warn",
	Error: "error",
	Fatal: "fatal",
}

var levelsStrings = map[string]Level{
	"trace": Trace,
	"debug": Debug,
	"info":  Info,
	"warn":  Warn,
	"error": Error,
	"fatal": Fatal,
}

func (l Level) String() string {
	return levelNames[l]
}

func parseLevel(s string) (Level, error) {
	level, ok := levelsStrings[s]
	if !ok {
		return InvalidLevel, levelError
	}
	return level, nil
}
