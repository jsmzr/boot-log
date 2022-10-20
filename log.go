package log

import (
	"fmt"
)

type LoggerConfig interface {
	Load() (Logger, error)
}

type Logger interface {
	Debug(msg string, paras ...interface{})
	Info(msg string, paras ...interface{})
	Warn(msg string, params ...interface{})
	Error(msg string, params ...interface{})
}

var adapters = make(map[string]LoggerConfig)
var instance Logger

func Register(name string, adapter LoggerConfig) {
	_, ok := adapters[name]
	if ok {
		panic(fmt.Errorf("adapter [%s] already register", name))
	}
	adapters[name] = adapter
}

func InitLogger(name string) error {
	adapter, ok := adapters[name]
	if !ok {
		return fmt.Errorf("not found adapter [%s]", name)
	}
	if instance != nil {
		fmt.Printf("Instance already init, will be [%s] replaced\n", name)
	}
	if newInstance, err := adapter.Load(); err != nil {
		return err
	} else {
		instance = newInstance
		return nil
	}
}

func Debug(msg string, params ...interface{}) {
	instance.Debug(msg, params...)
}

func Info(msg string, params ...interface{}) {
	instance.Info(msg, params...)
}

func Warn(msg string, params ...interface{}) {
	instance.Warn(msg, params...)
}

func Error(msg string, params ...interface{}) {
	instance.Error(msg, params...)
}
