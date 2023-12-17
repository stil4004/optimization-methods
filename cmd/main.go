package main

import (
	"Opt/invest"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	// // Creating the condition for backpack
	// var cond backpack.Condition
	// cond.CreateCondition()
	// cond.Solve()
	// cond.PrintTables()

	// Creating the condition for inveest
	var cond invest.Condition
	cond.Input()
	cond.Solve()
	cond.PrintTables()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<- sigChan
}