package ylog

import (
	"encoding/json"
	"strings"
)

type (
	Formatter interface {
		Format(r Record) (string, error)
	}

	LineFormatter struct {
	}

	JsonFormatter struct {
	}
)

func NewLineFormatter() Formatter {
	return &LineFormatter{}
}

func NewJsonFormatter() Formatter {
	return &JsonFormatter{}
}

func (l LineFormatter) Format(r Record) (string, error) {
	var builder strings.Builder

	builder.WriteString(r.Time.Format("2006-01-02 15:04:05"))
	builder.WriteString(" ")
	builder.WriteString(strings.ToUpper(r.Level.String()))
	builder.WriteString(" ")
	builder.WriteString(r.Msg)

	for k, v := range r.Fields {
		builder.WriteString(" ")
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(InterfaceToStr(v))
	}

	builder.WriteString("\n")

	return builder.String(), nil
}

func (j JsonFormatter) Format(r Record) (string, error) {
	newdata := make(Fields, 3+len(r.Fields))

	newdata["time"] = r.Time.Format("2006-01-02 15:04:05")
	newdata["level"] = r.Level.String()
	newdata["msg"] = r.Msg

	for k, v := range r.Fields {
		newdata[k] = v
	}

	str, _ := json.Marshal(newdata)

	return string(str) + "\n", nil
}
