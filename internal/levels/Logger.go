package levels

import (
	"fmt"
	"os"
	"time"
)

var logtime = time.Now().Format("02 01 2006 15:04:05 Mon")

var errorColor = "\033[31m"
var debugColor = "\033[97m"
var warningColor = "\033[35m"
var infoColor = "\033[34m"
var resetColor = "\033[0m"

type Logger struct {
	FilePath      string
	DebugOption   bool
	WarningOption bool
	ErrorOption   bool
	InfoOption    bool
}

func (thisLogger Logger) Warning(message string, console bool, data ...interface{}) {
	if thisLogger.WarningOption {
		result := fmt.Sprintf("---WARN--- "+"\n"+" %s ┃ %s %#v "+"\n"+"---WARN--- "+"\n\n"+" ", logtime, message, data)
		switch console {
		case true:
			fmt.Println(warningColor + result + resetColor)
		case false:
			file, err := os.OpenFile(thisLogger.FilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			_, err = fmt.Fprintf(file, result)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (thisLogger Logger) Info(message string, console bool, data ...interface{}) {
	if thisLogger.InfoOption {
		result := fmt.Sprintf("---INFO--- "+"\n"+" %s ┃ %s %#v "+"\n"+"---INFO--- "+"\n\n"+" ", logtime, message, data)
		switch console {
		case true:
			fmt.Println(infoColor + result + resetColor)
		case false:
			file, err := os.OpenFile(thisLogger.FilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			_, err = fmt.Fprintf(file, result)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
func (thisLogger Logger) Debug(message string, console bool, data ...interface{}) {
	if thisLogger.DebugOption {
		result := fmt.Sprintf("---DEBUG--- "+"\n"+" %s ┃ %s %#v "+"\n"+"---DEBUG--- "+"\n\n"+" ", logtime, message, data)
		switch console {
		case true:
			fmt.Println(debugColor + result + resetColor)
		case false:
			file, err := os.OpenFile(thisLogger.FilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			_, err = fmt.Fprintf(file, result)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
func (thisLogger Logger) Error(message string, console bool, data ...interface{}) {
	if thisLogger.ErrorOption {
		result := fmt.Sprintf("---ERROR--- "+"\n"+" %s ┃ %s %#v "+"\n"+"---ERROR--- "+"\n\n"+" ", logtime, message, data)
		switch console {
		case true:
			fmt.Println(errorColor + result + resetColor)
		case false:
			file, err := os.OpenFile(thisLogger.FilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			_, err = fmt.Fprintf(file, result)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
