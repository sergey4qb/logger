package main

import (
	logger "github.com/sergey4qb/logger/internal"
	"net"
)

func main() {
	logg := logger.NewMyLogger("./outputLog/log", 1, 1, 1)
	logg.Info("Test", false, net.Interface{
		Index:        0,
		MTU:          0,
		Name:         "",
		HardwareAddr: nil,
		Flags:        0,
	})
}
