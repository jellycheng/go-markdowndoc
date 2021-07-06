package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"markdowndoc/config"
	"io/ioutil"
	"net/http"
	"runtime"
)

// RecoveryWithWriter returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func RecoveryMiddle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			traceId := c.GetString("traceid")

			stack := append([]byte(err.(error).Error() + "\n"), stack(3)...)
			stack = bytes.ReplaceAll(stack, []byte("\n"), []byte("\\n"))
			stack = bytes.ReplaceAll(stack, []byte("\""), []byte("\\\""))
			stack = bytes.ReplaceAll(stack, []byte("\t"), []byte("\\t"))

			logrus.Errorf("%s %s \"panic recovered: %s\"", traceId, c.FullPath(), stack)

			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
				"code": config.ErrorDefault,
				"msg": config.GetMsg4Code(config.ErrorDefault),
				"trace_id": traceId,
				"data": struct {}{},
			})
		}
	}()
	c.Next()
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		_, _ = fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return []byte("???")
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return []byte("???")
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, []byte("/")); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, []byte(".")); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, []byte("·"), []byte("."), -1)
	return name
}