package ylog

import (
	"fmt"
	"strconv"
)

type Fields map[string]interface{}

func InterfaceToStr(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case uint:
		return strconv.Itoa(int(v.(uint)))
	case uint8:
		return strconv.Itoa(int(v.(uint8)))
	case int8:
		return strconv.Itoa(int(v.(int8)))
	case uint16:
		return strconv.Itoa(int(v.(uint16)))
	case int16:
		return strconv.Itoa(int(v.(int16)))
	case uint32:
		return strconv.Itoa(int(v.(uint32)))
	case int32:
		return strconv.Itoa(int(v.(int32)))
	case uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case []byte:
		return string(v.([]byte))
	case bool:
		return strconv.FormatBool(v.(bool))
	default:
		return fmt.Sprintf("%v", v)
	}
}
