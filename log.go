// implemente the httpRPC Logger interface
package api

import (
	"fmt"
	"time"
)

type ConsoleLogger struct{
}

func (lg ConsoleLogger) ReqId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (lg ConsoleLogger) Xput(logs []string) {
	for _, log := range logs {
		fmt.Println(log)
	}
}

func (lg ConsoleLogger) Trace(format string, args ...interface{}) {
	fmt.Printf("[TRACE] " + format + "\n", args...)
}
func (lg ConsoleLogger) Debug(format string, args ...interface{}){
	fmt.Printf("[DEBUG] " + format + "\n", args...)	
}
func (lg ConsoleLogger) Info(format string, args ...interface{}){
	fmt.Printf("[INFO] " + format + "\n", args...)
}
func (lg ConsoleLogger) Warn(format string, args ...interface{}) {
	fmt.Printf("[WARN] " + format + "\n", args...)
}
func (lg ConsoleLogger) Error(format string, args ...interface{}) {
	fmt.Printf("[ERROR] " + format + "\n", args...)
}
func (lg ConsoleLogger) Critical(format string, args ...interface{}){
	fmt.Printf("[CRITICAL] " + format + "\n", args...)
}	
