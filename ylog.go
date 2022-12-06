package ylog

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type (
	Logger interface {
		Info(msg string)
		Debug(msg string)
		Warn(msg string)
		Error(msg string)

		SetFormatter(f Formatter)
		SetOuter(o Outer)
		SetLevel(l Level)
		With(fields Fields) Logger
	}

	// 日志记录
	Record struct {
		Time   time.Time
		Msg    string // 原始消息
		Fmsg   string // 格式化后的消息内容
		Level  Level
		Fields Fields // 自定义字段
	}

	Log struct {
		f      Formatter
		o      Outer
		level  Level
		fields Fields // 公共字段
		mu     sync.Mutex
	}
)

var defaultLog Logger

func init() {
	defaultLog = New()
	defaultLog.SetLevel(InfoLevel)
}

func New() Logger {
	return &Log{
		f: NewLineFormatter(),
		o: NewStdout(),
	}
}

func DefaultLogger() Logger {
	return defaultLog
}

func Info(msg string) {
	defaultLog.Info(msg)
}
func Infof(format string, a ...interface{}) {
	defaultLog.Info(fmt.Sprintf(format, a...))
}
func Debug(msg string) {
	defaultLog.Debug(msg)
}
func Debugf(format string, a ...interface{}) {
	defaultLog.Debug(fmt.Sprintf(format, a...))
}
func Warn(msg string) {
	defaultLog.Warn(msg)
}
func Warnf(format string, a ...interface{}) {
	defaultLog.Warn(fmt.Sprintf(format, a...))
}
func Error(msg string) {
	defaultLog.Error(msg)
}
func Errorf(format string, a ...interface{}) {
	defaultLog.Error(fmt.Sprintf(format, a...))
}
func SetOuter(o Outer) {
	defaultLog.SetOuter(o)
}
func SetFormatter(f Formatter) {
	defaultLog.SetFormatter(f)
}
func SetLevel(l Level) {
	defaultLog.SetLevel(l)
}

func (l *Log) Info(msg string) {
	l.Write(msg, InfoLevel)
}
func (l *Log) Debug(msg string) {
	l.Write(msg, DebugLevel)
}
func (l *Log) Warn(msg string) {
	l.Write(msg, WarnLevel)
}
func (l *Log) Error(msg string) {
	l.Write(msg, ErrorLevel)
}

func (l *Log) SetFormatter(f Formatter) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.f = f
}
func (l *Log) GetFormatter() Formatter {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.f
}

func (l *Log) SetOuter(o Outer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.o = o
}
func (l *Log) GetOuter() Outer {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.o
}

func (l *Log) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&l.level), uint32(level))
}
func (l *Log) GetLevel() Level {
	return Level(atomic.LoadUint32((*uint32)(&l.level)))
}

func (l *Log) With(fs Fields) Logger {
	newdata := make(Fields, len(l.fields)+len(fs))
	for k, v := range l.GetFields() {
		newdata[k] = v
	}
	for k, v := range fs {
		newdata[k] = v
	}
	return &Log{
		f:      l.GetFormatter(),
		o:      l.GetOuter(),
		level:  l.GetLevel(),
		fields: newdata,
	}
}
func (l *Log) GetFields() Fields {
	return l.fields
}

func (l *Log) Write(msg string, level Level) error {

	if level >= l.GetLevel() {
		record := Record{
			Time:   time.Now(),
			Msg:    msg,
			Level:  level,
			Fields: l.GetFields(),
		}

		// 格式化
		fmsg, err := l.GetFormatter().Format(record)
		if err != nil {
			return fmt.Errorf("format record err: %w", err)
		}
		record.Fmsg = fmsg

		// 输出
		return l.GetOuter().Write(record)
	}
	return nil
}
