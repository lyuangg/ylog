package benchmarks

import (
	"github.com/lyuangg/ylog"
)

func newYlogDefaultLogger() ylog.Logger {
	eo := ylog.NewEmptyOut()
	log := ylog.New()
	log.SetOuter(eo)

	return log
}

func newYlogRoateLogger() ylog.Logger {
	log := ylog.New()
	fo := ylog.NewRotateFile("./testdata/ylog", "ylog-2006-01-02.log", 1)
	log.SetOuter(fo)

	return log
}

func newYlogJsonLogger() ylog.Logger {
	eo := ylog.NewEmptyOut()
	log := ylog.New()
	log.SetOuter(eo)
	log.SetFormatter(ylog.NewJsonFormatter())

	return log
}
