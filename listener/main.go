package main

import (
	"go-stater-listener/servicebus"
)

func main() {
	run()
}

func run() {
	// appinsight := appinsight.NewAppinsights()
	// logT := utils.TerminalLogger()
	// logM := bpLogCenter.NewLogCenter(appinsight, logT)

	s := servicebus.NewServiceBus()

	// go health.Health(logM)
	s.SubscriptionSuccess()
}
