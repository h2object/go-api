package api

import (
	"fmt"
	"sync"
	"github.com/h2object/rpc"
)

type Auth interface{
	rpc.PreRequest
}

type Logger interface{
	rpc.Logger
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{}) 
	Error(format string, args ...interface{}) 
	Critical(format string, args ...interface{})
}

var UserAgent = "Golang h2object/go-api package"

type H2Object struct{
	sync.RWMutex
	addr string
	conn *rpc.Client
}

func NewH2Object(host string, port int) *H2Object {
	connection := rpc.NewClient(rpc.H2OAnalyser{})
	h2o := &H2Object{
		addr: fmt.Sprintf("%s:%d", host, port),	
		conn: connection,
	}
	return h2o
}


