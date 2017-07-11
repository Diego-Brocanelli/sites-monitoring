package main

import (
	"fmt"
	"os"
	"sites-monitoring/log"
	"sites-monitoring/monitor"
	"sites-monitoring/options"
)

const lineSeparator = "------------------------------------------------------------------"
const chooseCloseSystem = 0
const chooseStarMonitoring = 1
const chooseShowLogs = 2

func main() {
	options.ShowOptions()

	command := options.GetCommand()

	switch command {
	case chooseCloseSystem:
		fmt.Println("Thank you for use the monitoring system...goodbye :)")
		os.Exit(0)
		break
	case chooseStarMonitoring:
		fmt.Println()
		fmt.Println(lineSeparator)
		fmt.Println("Start monitoring")

		quantity := options.GetQuantity()

		monitor.Start(quantity)

		fmt.Println(lineSeparator)
		fmt.Println("Finished monitoring")
		fmt.Println("")
		os.Exit(0)
		break
	case chooseShowLogs:
		fmt.Println("Show logs")
		fmt.Println(lineSeparator)

		log.Show()

		fmt.Println(lineSeparator)
		os.Exit(0)
		break
	default:
		fmt.Println("Command not suported...")
		os.Exit(-1)
		break
	}
}
