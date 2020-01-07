package gocommon

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var currLogLevel = LogLevel_INFO

func LOGprefix(level LogLevel, skip int, args ...interface{}) string {
	if level < currLogLevel {
		return ""
	}
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(skip)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}

	prefix := fmt.Sprintf("%s:%d:%s: %s", filename, line, funcname, level)
	return prefix
}
