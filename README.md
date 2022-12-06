# ylog

一个简单易用的 go 日志库。

## features

- 支持日志级别: debug, info, warning, error
- 支持按日期旋转文件
- 支持 json 格式
- 支持自定义格式和输出
- 代码测试覆盖率 100%

## install

```
go get -u github.com/lyuangg/ylog
```

## usage

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

## test

```bash
go test -cover
```

output:

```bash
PASS
coverage: 100.0% of statements
ok      github.com/lyuangg/ylog 0.109s
```