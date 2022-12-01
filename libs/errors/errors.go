// Package errors 错误封装
package errors

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type (
	// stack 堆栈信息
	stack struct {
		str string
		pc  [32]uintptr
		n   int
	}
	// stackError 带有堆栈信息的错误
	stackError struct {
		stack
		code int
		info string
		msg  string
	}
)

// getStack 获取堆栈指针，但不进一步处理
func (s *stack) getStack() {
	s.n = runtime.Callers(3, s.pc[:])
}

// stackInfo 获取堆栈详细信息
func (s *stack) stackInfo() string {
	if s.str == "" {
		s.str = "\n错误堆栈:"
		for i := 0; i < s.n; i++ {
			fn := runtime.FuncForPC(s.pc[i] - 1)
			if fn == nil {
				s.str += "\n??? ???:???"
				continue
			}
			functionName := fn.Name()
			if strings.Index(functionName, "runtime.") != -1 {
				continue
			}
			if strings.Index(functionName, ".ServeHTTP") != -1 {
				continue
			}
			if strings.Index(functionName, ".serve") != -1 {
				continue
			}
			s.str += "\n" + functionName + " "
			file, line := fn.FileLine(s.pc[i] - 1)
			s.str += file + ":" + strconv.Itoa(line)
		}
	}
	return s.str
}

// Error 返回不带堆栈信息的错误
func (s *stackError) Error() string {
	return s.msg
}

// Format 写入堆栈信息和错误信息
func (s *stackError) Format(f fmt.State, _ rune) {
	_, _ = f.Write([]byte(s.msg + s.stackInfo()))
}

// New 生成新的带堆栈信息的错误，或为给定错误附带堆栈信息
func New(v ...interface{}) error {
	s := &stackError{info: "未知错误", msg: "未知错误", code: 0}
	getStack := true
	if len(v) > 0 {
		switch v[0].(type) {
		case nil:
			return nil
		case string, error:
			switch v[0].(type) {
			case string:
				s.msg = v[0].(string)
				s.info = s.msg
			case error:
				s1, ok := v[0].(*stackError)
				if ok {
					s = s1
					getStack = false
				} else {
					s.msg = v[0].(error).Error()
					s.info = s.msg
				}
			}
			if len(v) > 1 {
				switch v[1].(type) {
				case string:
					s.info = v[1].(string)
				}
			}
		}
	}
	if getStack {
		s.getStack()
	}
	return s
}

func GetInfo(err error) string {
	e, ok := err.(*stackError)
	if ok {
		return e.info
	}
	return err.Error()
}
