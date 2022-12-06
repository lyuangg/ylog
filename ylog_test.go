package ylog

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	TestOuter struct {
		Msg string
	}

	TestFormatter struct {
	}

	TestErrFormatter struct {
	}
)

func (t *TestOuter) Write(r Record) error {
	t.Msg = r.Fmsg
	return nil
}

func (t TestFormatter) Format(r Record) (string, error) {
	return r.Msg, nil
}
func (t TestErrFormatter) Format(r Record) (string, error) {
	return "", errors.New("err format")
}

func TestLoggerDataRace(t *testing.T) {
	w := &sync.WaitGroup{}
	log := New()
	for i := 0; i < 5; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			log.SetOuter(NewEmptyOut())
			log.SetLevel(DebugLevel)
			log.SetFormatter(NewJsonFormatter())
			log.With(Fields{"hello": "good"}).Info("with info..")
		}()
	}
	w.Wait()
}

func TestSetOuter(t *testing.T) {
	o := &TestOuter{}
	SetOuter(o)
	Info("hello world")
	assert.Contains(t, o.Msg, "hello world")
}

func TestSetLevel(t *testing.T) {
	o := &TestOuter{}
	SetOuter(o)
	SetLevel(DebugLevel)
	Info("hello world")
	assert.Contains(t, o.Msg, "hello world")
	assert.Contains(t, o.Msg, "INFO")

	o.Msg = ""
	SetLevel(WarnLevel)
	Info("hello world")
	assert.Empty(t, o.Msg)
}

func TestSetFormatter(t *testing.T) {
	o := &TestOuter{}
	f := TestFormatter{}
	SetOuter(o)
	SetLevel(DebugLevel)
	SetFormatter(f)
	Info("hello world")
	assert.Equal(t, o.Msg, "hello world")
}

func TestLevelWrite(t *testing.T) {
	o := &TestOuter{}
	SetOuter(o)
	SetLevel(DebugLevel)
	SetFormatter(NewLineFormatter())

	Infof("%s", "hello")
	assert.Contains(t, o.Msg, "hello")
	o.Msg = ""

	Debug("debug info")
	assert.Contains(t, o.Msg, "debug info")
	o.Msg = ""
	Debugf("debug info:%d,%d", 2, 3)
	assert.Contains(t, o.Msg, "debug info:2,3")
	o.Msg = ""

	Warn("warn str")
	assert.Contains(t, o.Msg, "warn str")
	o.Msg = ""
	Warnf("warn str:%d,%d", 2, 3)
	assert.Contains(t, o.Msg, "warn str:2,3")
	o.Msg = ""

	Error("error str")
	assert.Contains(t, o.Msg, "error str")
	o.Msg = ""
	Errorf("error str:%d,%d", 2, 3)
	assert.Contains(t, o.Msg, "error str:2,3")
	o.Msg = ""

	DefaultLogger().Info("info hello")
	assert.Contains(t, o.Msg, "info hello")
}

func TestWith(t *testing.T) {
	o := &TestOuter{}
	SetOuter(o)

	With(Fields{"testkey1": "testval1", "testkey2": "testval2"}).With(Fields{"testkey3": "testval3"}).Info("hello world")

	assert.Contains(t, o.Msg, "testkey1")
	assert.Contains(t, o.Msg, "testval2")
	assert.Contains(t, o.Msg, "hello world")
}

func TestWriteErr(t *testing.T) {
	log := &Log{}
	log.SetFormatter(TestErrFormatter{})

	err := log.Write("hello", DebugLevel)
	assert.NotNil(t, err)
}
