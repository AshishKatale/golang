package main

import (
	"learn/features"
)

func main() {

	// fmt.Println(features.GetPrimesInRangeParallel(0, 250000, 10))

	// var till int
	// var jobs int
	//
	// flag.IntVar(&till, "t", 100, "find primes till")
	// flag.IntVar(&jobs, "j", 5, "no of goroutines to learn")
	// showHelp := flag.Bool("h", false, "help text")
	// flag.Parse()
	//
	// if *showHelp {
	// 	flag.CommandLine.Usage()
	// 	os.Exit(0)
	// }
	//
	// p := features.GetPrimesInRangeParallel(0, till, jobs)
	// fmt.Println(p)

	// getFib := features.FibonacciClosure()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(getFib())
	// }

	// for i := uint(0); i < 10; i++ {
	// 	fmt.Println(features.Fibonacci(i))
	// }
	// features.RunChannel()
	// features.RunMutex()
	features.FetchAndParse()

}
