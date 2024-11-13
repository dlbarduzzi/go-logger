package main

import "github.com/dlbarduzzi/showcase/internal/logging"

func main() {
	logger := logging.NewLogger("dev", "debug")
	logger.Info("Hello from dev!")

	logger2 := logging.NewLogger("prod", "debug")
	logger2.Info("Hello from prod!")
}
