package helpers

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func Empty(val interface{}) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
	}
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}

// AgentToOs 从 agent 判断系统
func AgentToOs(agent string) string {
	// 判断是否包含 'win' 并匹配特定的 Windows 版本
	if strings.Contains(strings.ToLower(agent), "win") {
		if matched, _ := regexp.MatchString(`nt 6.1`, agent); matched {
			return "Windows 7"
		}
		if matched, _ := regexp.MatchString(`nt 6.2`, agent); matched {
			return "Windows 8"
		}
		if matched, _ := regexp.MatchString(`nt 10.0`, agent); matched {
			return "Windows 10"
		}
		if matched, _ := regexp.MatchString(`nt 11.0`, agent); matched {
			return "Windows 11"
		}
		if matched, _ := regexp.MatchString(`nt 5.1`, agent); matched {
			return "Windows XP"
		}
	}

	// 判断是否包含 'linux'
	if strings.Contains(strings.ToLower(agent), "linux") {
		return "Linux"
	}

	// 判断是否包含 'mac'
	if strings.Contains(strings.ToLower(agent), "mac") {
		return "Mac"
	}

	// 默认返回 Unknown
	return "Unknown"
}

// AgentToBrowser 从 agent 判断浏览器
func AgentToBrowser(agent string) string {
	// 判断浏览器类型
	if strings.Contains(strings.ToLower(agent), "msie") {
		return "MSIE"
	}
	if strings.Contains(strings.ToLower(agent), "edg") {
		return "Edge"
	}
	if strings.Contains(strings.ToLower(agent), "chrome") {
		return "Chrome"
	}
	if strings.Contains(strings.ToLower(agent), "firefox") {
		return "Firefox"
	}
	if strings.Contains(strings.ToLower(agent), "safari") {
		return "Safari"
	}
	if strings.Contains(strings.ToLower(agent), "opera") {
		return "Opera"
	}

	// 默认返回 Unknown
	return "Unknown"
}
