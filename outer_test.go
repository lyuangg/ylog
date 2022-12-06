package ylog

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type (
	TestLogger struct {
		Logger
	}
)

func (t *TestLogger) GetFormatter() Formatter {
	return LineFormatter{}
}

func TestStdouter(t *testing.T) {
	sout := NewStdout()
	assert.NotNil(t, sout)

	stdout := &StdOut{o: io.Discard}

	r := Record{
		Time:   time.Now(),
		Level:  InfoLevel,
		Msg:    "hello",
		Fields: Fields{"good": "test"},
	}
	var err error
	r.Fmsg, err = NewLineFormatter().Format(r)
	err = stdout.Write(r)
	assert.Nil(t, err)
}

func TestNewRoate(t *testing.T) {
	f := NewRotateFile("./testdata/", "new-2006-01-02", 0)
	assert.NotNil(t, f)
}

func TestNameFormat(t *testing.T) {
	f := &RotateFileOut{
		path:       "./testdata",
		nameFormat: "testlog-2006-01-02",
		num:        0,
	}

	tm := time.Now()
	resultName := "testlog-" + tm.Format("2006-01-02")
	assert.Contains(t, f.fileNameByTime(tm), resultName)
	assert.Contains(t, f.fileNameByTime(tm), "./testdata")

}

func TestRemoveDir(t *testing.T) {
	p := "./testdata/testdir"
	f := &RotateFileOut{
		path:       p,
		nameFormat: "testlog-2006-01-02",
		num:        0,
	}
	err := f.createDir()
	assert.Nil(t, err)

	assert.DirExists(t, p)
	os.Remove(p)
}

func TestGlobFiles(t *testing.T) {
	p := "./testdata/testglob"
	f := &RotateFileOut{
		path:       p,
		nameFormat: "testlog-2006-01-02",
		num:        0,
	}
	err := f.createDir()
	assert.Nil(t, err)

	filelist := []string{
		"/testlog-2022-01-01",
		"/testlog-2022-01-02",
		"/testlog-2022-01-03",
	}
	for _, v := range filelist {
		tf, _ := os.Create(p + v)
		tf.Close()
	}

	fs := f.globFiles()
	assert.Len(t, fs, len(filelist))
	os.RemoveAll(p)
}

func TestDelFiles(t *testing.T) {
	p := "./testdata/testdel"
	nf := "testlog-2006-01-02"

	f := &RotateFileOut{
		path:       p,
		nameFormat: nf,
		num:        1,
	}
	err := f.createDir()
	assert.Nil(t, err)

	filelist := []string{
		"/testlog-2022-01-01",
		"/testlog-2022-01-02",
		"/testlog-2022-01-03",
	}
	for _, v := range filelist {
		tf, _ := os.Create(p + v)
		tf.Close()
	}

	err = f.delFiles()
	assert.Nil(t, err, "delfile")

	fs := f.globFiles()
	assert.Len(t, fs, 1)
	os.RemoveAll(p)
}

func TestRotateFileOut(t *testing.T) {
	fout := &RotateFileOut{
		path:       "./testdata",
		nameFormat: "testroutefile-2006-01-02.log",
		num:        1,
	}

	r := Record{
		Time:  time.Now(),
		Level: InfoLevel,
		Msg:   "hello",
	}
	var err error
	r.Fmsg, err = NewLineFormatter().Format(r)
	err = fout.Write(r)
	err = fout.Write(r)
	defer os.Remove(fout.openName)

	assert.Nil(t, err)
	assert.FileExists(t, fout.openName)

	r.Time, _ = time.Parse("2006-01-02", "2022-01-01")
	fout.Write(r)

	os.Remove(fout.openName)
}

func TestWrite(t *testing.T) {
	fout := &RotateFileOut{
		path:       "./testdata",
		nameFormat: "testwritefile-2006-01-02.log",
		num:        1,
	}

	r := Record{
		Time:  time.Now(),
		Level: InfoLevel,
		Msg:   "hello",
	}
	var err error
	r.Fmsg, err = NewLineFormatter().Format(r)

	path := fout.fileNameByTime(r.Time)
	fd, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0555)
	defer os.Remove(path)
	defer fd.Close()

	err = fout.Write(r)
	assert.NotNil(t, err)

	fout.openName = fout.fileNameByTime(r.Time)
	err = fout.Write(r)
	assert.NotNil(t, err)
}
