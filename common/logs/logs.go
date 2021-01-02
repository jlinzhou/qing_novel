package logs

import (
	"bytes"
	"fmt"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)


// line number hook for log the call context,
type lineHook struct {
	Field  string
	// skip为遍历调用栈开始的索引位置
	Skip   int
	levels []log.Level
}

// Levels implement levels
func (hook lineHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire implement fire
func (hook lineHook) Fire(entry *log.Entry) error {
	entry.Data[hook.Field] = findCaller(hook.Skip)
	return nil
}

func findCaller(skip int) string {
	file := ""
	line := 0
	//var pc uintptr
	// 遍历调用栈的最大索引为第11层.
	for i := 0; i < 11; i++ {
		file, line, _ = getCaller(skip + i)
		// 过滤掉所有logrus包，即可得到生成代码信息
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}
	n := 0

	// 获取包名
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, pc
}
// MyFormatter 自定义 formatter
type MyFormatter struct {
}
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func (mf *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timeParse:= entry.Time.Format("2006-01-02 15:04:05.000")
	threadId:=getGID()
	var path interface{}
	if _,ok:=entry.Data["path"];ok{
		path=entry.Data["path"]
	}
	b.WriteString(fmt.Sprintf("%s %s [%d] %s [%s]\n", timeParse,entry.Level, threadId, entry.Message,path))
	return b.Bytes(), nil
}

func New(filePath string,maxsize int,maxBackups int,maxAge int,compress bool){
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxsize,  // Max megabytes before log is rotated
		MaxBackups: maxBackups, // Max number of old log files to keep
		MaxAge:     maxAge, // Max number of days to retain log files
		Compress:   compress,
	}

	formatter := &MyFormatter{
	}
	log.SetFormatter(formatter)
	logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
	log.SetOutput(logMultiWriter)
	hook:=lineHook{Skip:3,Field:"path"}
	log.AddHook(&hook)
}