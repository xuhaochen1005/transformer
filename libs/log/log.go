// Package log 日志记录模块
package log

import (
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"transformer_mz/libs/config"
	"transformer_mz/libs/errors"
)

var Logger = logrus.New()

func Trace(v ...interface{}) {
	Logger.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
	Logger.Tracef(format, v...)
}

func Traceln(v ...interface{}) {
	Logger.Traceln(v...)
}

func Debug(v ...interface{}) {
	Logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	Logger.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	Logger.Debugln(v...)
}

func Info(v ...interface{}) {
	Logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Logger.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	Logger.Infoln(v...)
}

func Print(v ...interface{}) {
	Logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

func Println(v ...interface{}) {
	Logger.Println(v...)
}

func Warn(v ...interface{}) {
	Logger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	Logger.Warnf(format, v...)
}

func Warnln(v ...interface{}) {
	Logger.Warnln(v...)
}

func Error(v ...interface{}) {
	Logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	Logger.Errorln(v...)
}

func Fatal(v ...interface{}) {
	Logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	Logger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	Logger.Fatalln(v...)
}

func Panic(v ...interface{}) {
	Logger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	Logger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	Logger.Panicln(v...)
}

type ConfigWrap struct {
	Level  logrus.Level `yaml:"level"`
	File   string       `yaml:"file"`
	Rotate int          `yaml:"rotate"`
	Keep   int          `yaml:"keep"`
}

type Config struct {
	ConfigWrap `yaml:"log"`
}

// InitLog 初始化日志模块，日志内容除了输出到屏幕，也会以文件的形式保存并自动轮转和删除
func InitLog() error {
	logConfig := &Config{}
	err := config.Load(logConfig)
	if err != nil {
		return err
	}
	Logger.SetLevel(logConfig.Level)
	Logger.SetFormatter(&logrus.TextFormatter{DisableColors: true, DisableQuote: true})
	writer, err := rotatelogs.New(
		logConfig.File+".%Y%m%d",
		rotatelogs.WithLinkName(logConfig.File),
		rotatelogs.WithRotationTime(time.Duration(logConfig.Rotate)*24*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(logConfig.Keep)*24*time.Hour))
	if err != nil {
		return errors.New(err)
	}
	Logger.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.TraceLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
		logrus.FatalLevel: writer,
	}, nil))
	return nil
}
