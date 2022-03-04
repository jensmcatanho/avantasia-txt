package main

import (
	"github.com/jensmcatanho/avantasia-txt/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	avantasiaTxt := cmd.NewAvantasiaTxt(logger)
	avantasiaTxt.Run()
}
