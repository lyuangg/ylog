# ylog

一个简单易用的 go 日志库。

## Features

- 支持日志级别: debug, info, warning, error
- 支持按日期旋转文件
- 支持 json 格式
- 支持自定义格式和输出
- 代码测试覆盖率 100%
- 线程安全的
- 超越 `sirupsen/logrus` 日志库的性能

## Install

```
go get -u github.com/lyuangg/ylog
```

## Usage

```go
// 基本使用
ylog.Info("hello")

ylog.Infof("hello %d", 1)

ylog.Debug("debug msg..")

ylog.Errorf("err msg: %s, %v", "err", "msg")

// json
ylog.SetFormatter(NewJsonFormatter())

// 自定义字段
ylog.With(Fields{
    "requestid" : "123456",
    "traceid" : 5555,
}).Info("hello ylog")

// 按日期旋转文件
outer := ylog.NewRotateFile("./testdata", "go-2006-01-02.log", 7)
ylog.SetOuter(outer)

// 设置日志级别
ylog.SetLevel(ylog.ParseLevel("debug"))
ylog.SetLevel(ylog.DebugLevel)

// new logger
newlog := ylog.New()
newlog.SetLevel(ylog.WarnLevel)

// 自定义格式和输出, 实现下面两个接口，调用 SetFormatter，SetOuter 方法
type (
    Formatter interface {
		Format(r Record) (string, error)
	}

    Outer interface {
		Write(r Record) error
	}
)
```

## Test

```bash
go test -cover
```

output:

```bash
PASS
coverage: 100.0% of statements
ok      github.com/lyuangg/ylog 0.109s
```

## Benchmark

```bash
go test -bench=. -benchtime=3s -benchmem -run=none
```

```bash
goos: darwin
goarch: amd64
pkg: github.com/lyuangg/ylog/benchmarks
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkDefaultLogger/logrus-8                  1528854              2652 ns/op             576 B/op         17 allocs/op
BenchmarkDefaultLogger/ylog-8                   11159875               307.4 ns/op           276 B/op          6 allocs/op
BenchmarkRoateLogger/logrus-8                     250870             13076 ns/op             945 B/op         21 allocs/op
BenchmarkRoateLogger/ylog-8                       622275              5385 ns/op             292 B/op          8 allocs/op
BenchmarkWithField/logrus-8                      1000000              3094 ns/op             881 B/op         18 allocs/op
BenchmarkWithField/ylog-8                        8655415               387.6 ns/op           500 B/op          7 allocs/op
BenchmarkJsonFormat/logrus-8                      879865              3747 ns/op            1682 B/op         29 allocs/op
BenchmarkJsonFormat/ylog-8                       3995971               999.0 ns/op          1585 B/op         25 allocs/op
PASS
ok      github.com/lyuangg/ylog/benchmarks      32.649s
```