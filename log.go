package mylog

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags)

func IfErrDie(v error) {
	if v != nil {
		logger.Output(2, fmt.Sprintf("%v", v))
		panic(v)
	}
}
func IfErrPrt(v error) bool {
	if v != nil {
		logger.Output(2, fmt.Sprintf("%v", v))
		return true

	}
	return false
}
func Logf(f string, v ...interface{}) {

	logger.Output(2, fmt.Sprintf(f, v...))

}
