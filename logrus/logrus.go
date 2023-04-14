package logrus

import (
	"fmt"
	"os"
)

func Info(msg ...interface{}) {
	fmt.Println(msg...)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println()
}

func Debug(msg ...interface{}) {
	fmt.Println(msg...)
}

func Debugf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println()
}

func Error(msg ...interface{}) {
	fmt.Printf("ERROR: %s", msg...)
}

func Errorf(msg string, args ...interface{}) {
	imsg := fmt.Sprintf(msg, args...)
	fmt.Printf("ERROR: %s", imsg)
	fmt.Println()
}

func Warn(msg ...interface{}) {
	fmt.Printf("WARN: %s", msg...)
}

func Warnf(msg string, args ...interface{}) {
	imsg := fmt.Sprintf(msg, args...)
	fmt.Printf("WARN: %s", imsg)
	fmt.Println()
}

func Fatal(msg ...interface{}) {
	fmt.Println(msg...)
	os.Exit(1)
}

func Fatalf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println()
	os.Exit(1)
}

func Trace(msg ...interface{}) {
	fmt.Println(msg...)
}

func Tracef(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println()
}
