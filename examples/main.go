package main

import (
	"fmt"

	"github.com/davidgaspardev/golog"
)

func main() {
	golog.Info("Intelmes", "Intelmes is a manufacturing executation (MES) system to Intelbras")
	golog.Service("Intelmes", "Machines modbus started")
	golog.System("Intelmes", "Routes created")
	golog.Error("Intelmes", fmt.Errorf("Failed to generate OEE"))
	golog.Warn("Intelmes", "Unable to add resource")
}
