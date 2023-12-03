package main

import (
	"Opt/backpack"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	// Creating the condition
	var cond backpack.Condition
	cond.CreateCondition()
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