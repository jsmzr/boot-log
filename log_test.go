package log

import (
	"errors"
	"testing"
)

type FmtLoggerConfig struct{}
type ErrorLoggerConfig struct{}

type FmtLogger struct{}

func (f *FmtLoggerConfig) Load() (Logger, error) {
	return &FmtLogger{}, nil
}

func (e *ErrorLoggerConfig) Load() (Logger, error) {
	return nil, errors.New("load error")
}

func (f *FmtLogger) Debug(msg string, params ...interface{}) {

}
func (f *FmtLogger) Info(msg string, params ...interface{}) {

}
func (f *FmtLogger) Warn(msg string, params ...interface{}) {

}
func (f *FmtLogger) Error(msg string, params ...interface{}) {

}

func TestRegister(t *testing.T) {
	adapters = make(map[string]LoggerConfig)
	Register("test", &FmtLoggerConfig{})
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("should panic: adapter [test] already register")
		}
	}()
	Register("test", &FmtLoggerConfig{})
}

func TestInitLogger(t *testing.T) {
	adapters = make(map[string]LoggerConfig)
	instance = nil
	Register("test", &FmtLoggerConfig{})
	if err := InitLogger("test1"); err == nil {
		t.Fatal("should be not found test1 adapter")
	}

	if err := InitLogger("test"); err != nil {
		t.Fatal("should be found test adapter")
	}
	if instance == nil {
		t.Fatal("instance should be not nil")
	}
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Register("error", &ErrorLoggerConfig{})
	if err := InitLogger("error"); err == nil {
		t.Fatal("load adapter error should be error")
	}

}
