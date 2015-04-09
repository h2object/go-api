// implemente the httpRPC Logger interface
package api

import (
	"time"
	"strconv"
	"github.com/h2object/log"
)

// type Logger interface {
// 	ReqId() string
// 	Xput(logs []string)
// }

type Logger struct {
	*log.Logger
}

func NewLogger(lg *log.Logger) *Logger {
	return &Logger{lg}
}

func (lg *Logger) ReqId() string {
	return strconv.FormatInt(time.Now().UnixNano(), 16)
}

func (lg *Logger) Xput(logs []string) {
	for _, log := range logs {
		lg.Debug(log)
	}
}
