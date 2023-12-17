package main

import (
	"Opt/backpack"
	"Opt/invest"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func menu(){
	chose := 0
	for{
		fmt.Print("1. Backpack\n2. Invest\n")
		fmt.Scan(&chose)
		switch (chose){
		case 1:
			// Creating the condition for backpack
			var cond backpack.Condition
			cond.CreateCondition()
			cond.Solve()
			cond.PrintTables()
		case 2:
			// Creating the condition for inveest
			var cond invest.Condition
			cond.Input()
			cond.Solve()
			cond.PrintTables()
		default:
			fmt.Println("Wrong variant!")
		}
	}

}

func main() {
	args := os.Args[1:]
	if len(args) == 0{
		menu()
	}

	for _, ar := range args{
		switch (ar){
		case "-backpack":
			// Creating the condition for backpack
			var cond backpack.Condition
			cond.CreateCondition()
			cond.Solve()
			cond.PrintTables()
		case "-invest":
			// Creating the condition for inveest
			var cond invest.Condition
			cond.Input()
			cond.Solve()
			cond.PrintTables()
		default:
			fmt.Println("Wrong variant!")
		}
	}

	// // Creating the condition for backpack
	// var cond backpack.Condition
	// cond.CreateCondition()
	// cond.Solve()
	// cond.PrintTables()

	// Creating the condition for inveest
	// var cond invest.Condition
	// cond.Input()
	// cond.Solve()
	// cond.PrintTables()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<- sigChan
}