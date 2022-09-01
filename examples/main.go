package main

import (
	"fmt"

	"github.com/davidgaspardev/golog"
)

func main() {
	golog.Info("Intelmes", "Intelmes is a manufacturing executation (MES) system to Intelbras")
	golog.Io("Intelmes", "MODBUS request received")
	golog.Service("Intelmes", "Machines modbus started")
	golog.System("Intelmes", "Routes created")
	golog.Error("Intelmes", fmt.Errorf("Failed to generate OEE"))
	golog.Warn("Intelmes", "Unable to add resource")

	log := golog.NewLogTag("Database")
	log.Info("Create", "MES_resources table")
	log.Io("ListeningPlc", "pulse received")
	log.Service("Clean", "remove last data in MES_resources")
	log.System("Run", "start setup script")
	log.Error("GetItems", fmt.Errorf("columns not exists"))
	log.Warn("GetUserById", "id is only number")

	fmt.Println("\n------ Email Service Exmaple (README) ------")
	SendEmail("Hello World")
}

var log = golog.NewLogTag("Email Service")

func SendEmail(message string) {
	log.Service("SendEmail", "sending email with this content: "+message)
	// Code here
}
