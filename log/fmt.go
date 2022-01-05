package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
	"runtime"
)

type DefaultLog struct {
	log.Logger
}

func (d *DefaultLog) Fatal(i interface{}) {
	panic(i)
}

func (d *DefaultLog) Error(err error) {
	fmt.Printf("error: %v\n",err)
}

func (d *DefaultLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n",args...)
}

func (d *DefaultLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n",args...)
}



func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	})
}