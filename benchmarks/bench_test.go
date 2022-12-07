package benchmarks

import (
	"os"
	"testing"

	"github.com/lyuangg/ylog"
	"github.com/sirupsen/logrus"
)

func BenchmarkDefaultLogger(b *testing.B) {
	msg := "hello world,hello world,hello world,hello world"
	b.Run("logrus", func(b *testing.B) {
		logger := newLogrusDefaultLogger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})

	b.Run("ylog", func(b *testing.B) {
		logger := newYlogDefaultLogger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})
}

func BenchmarkRoateLogger(b *testing.B) {
	msg := "hello world,hello world,hello world,hello world"
	b.Run("logrus", func(b *testing.B) {
		logger := newLogrusRoateLogger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})
	b.Run("ylog", func(b *testing.B) {
		logger := newYlogRoateLogger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})

	defer os.RemoveAll("./testdata/logrus")
	defer os.RemoveAll("./testdata/ylog")
}

func BenchmarkWithField(b *testing.B) {
	msg := "hello world,hello world,hello world,hello world"
	b.Run("logrus", func(b *testing.B) {
		logger := newLogrusDefaultLogger().WithFields(logrus.Fields{
			"testkey1": "testval1",
			"testkey2": 2,
		})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})

	b.Run("ylog", func(b *testing.B) {
		logger := newYlogDefaultLogger().With(ylog.Fields{
			"testkey1": "testval1",
			"testkey2": 2,
		})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})
}

func BenchmarkJsonFormat(b *testing.B) {
	msg := "hello world,hello world,hello world,hello world"
	b.Run("logrus", func(b *testing.B) {
		logger := newLogrusJsonLogger().WithFields(logrus.Fields{
			"testkey1": "testval1",
			"testkey2": 2,
		})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})

	b.Run("ylog", func(b *testing.B) {
		logger := newYlogJsonLogger().With(ylog.Fields{
			"testkey1": "testval1",
			"testkey2": 2,
		})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})
}
