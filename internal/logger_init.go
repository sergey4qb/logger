package logger

import (
	"github.com/joho/godotenv"
	"github.com/sergey4qb/logger/internal/levels"
	"os"
	"strconv"
)

func NewMyLogger(f string, warning int, error int, info int) *levels.Logger {
	_, err := checkFile(f)
	if err != nil {
		panic(err)
	}

	return &levels.Logger{
		FilePath:      f,
		DebugOption:   getDebugEnv(),
		WarningOption: IntToBool(warning),
		ErrorOption:   IntToBool(error),
		InfoOption:    IntToBool(info),
	}
}

func IntToBool(i int) bool {
	var b bool = i != 0
	return b
}
func getDebugEnv() bool {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		panic(err)
	}
	return debug
}

func checkFile(filename string) (os.FileInfo, error) {
	file, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
	}
	return file, nil
}
