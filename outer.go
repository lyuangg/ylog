package ylog

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type (
	Outer interface {
		Write(r Record) error
	}

	// 标准输出
	StdOut struct {
		o  io.Writer
		mu sync.Mutex
	}

	// 空输出，测试时使用
	EmptyOut struct {
	}

	// 文件输出
	RotateFileOut struct {
		path       string
		nameFormat string
		num        int
		mu         sync.Mutex
		openName   string
		openFd     *os.File
	}
)

func NewStdout() Outer {
	return &StdOut{o: os.Stdout}
}

func NewEmptyOut() Outer {
	return &EmptyOut{}
}

func NewRotateFile(path string, nameformat string, num int) Outer {
	return &RotateFileOut{
		path:       path,
		nameFormat: nameformat,
		num:        num,
	}
}

func (e *EmptyOut) Write(r Record) error {
	return nil
}

func (s *StdOut) Write(r Record) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := s.o.Write([]byte(r.Fmsg))
	return err
}

func (ro *RotateFileOut) Write(r Record) error {

	ro.mu.Lock()
	defer ro.mu.Unlock()

	fpath := ro.fileNameByTime(r.Time)

	if fpath != ro.openName {
		if ro.openFd != nil {
			ro.openFd.Close()
			ro.openFd = nil
		}

		// 创建目录
		ro.createDir()

		// 打开日志文件
		fd, err := os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}

		ro.openName = fpath
		ro.openFd = fd

		// 检查文件数量, 删除多余的文件
		ro.delFiles()
	}

	// 写文件
	_, err := ro.openFd.WriteString(r.Fmsg)
	if err != nil {
		ro.openName = ""
		return err
	}

	return nil
}

// 根据时间获取文件名
func (ro *RotateFileOut) fileNameByTime(t time.Time) string {
	fname := t.Format(ro.nameFormat)
	fullpath := strings.TrimRight(ro.path, "/") + "/" + fname

	return fullpath
}

func (ro *RotateFileOut) createDir() error {
	if _, err := os.Stat(ro.path); err != nil {
		return os.MkdirAll(ro.path, 0755)
	}
	return nil
}

// 匹配路径下的所有文件
func (ro *RotateFileOut) globFiles() []string {
	re, _ := regexp.Compile("[0-9]+")
	namePattern := re.ReplaceAllString(ro.nameFormat, "*")
	pathPattern := strings.TrimRight(ro.path, "/") + "/" + namePattern
	files, _ := filepath.Glob(pathPattern)
	return files
}

// 删除多的文件
func (ro *RotateFileOut) delFiles() error {
	files := ro.globFiles()
	if len(files) > ro.num {
		sort.Strings(files)
		delFiles := files[0 : len(files)-ro.num]
		for _, f := range delFiles {
			os.Remove(f)
		}
	}
	return nil
}
