package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	fmt.Print("Performance measurement with make()")
}
func main() {
	// initialize()
	// withMake()
	withMake()
	// withOutMake()
}

func initialize() {
	var test int
	sports := make([]string, 4, 6)
	fmt.Print("Performance measurement with make()", test, sports)
}
func withMake() {
	start := time.Now()
	fmt.Println("Start time:", start)
	sports := make([]string, 8)
	sports = []string{"Cricket", "Football", "Running", "Badminton", "Baseball", "Hockey"}
	sports = append(sports, "Volleyball")
	sports = append(sports, "Swimming")
	fmt.Println("My favourite sports:")
	for index, sport := range sports {
		fmt.Printf("%d.\t%s\n", index+1, sport)
	}
	sports = append(sports[:3], sports[4:5]...)
	fmt.Println("My most favourite sports:")
	for index, sport := range sports {
		fmt.Printf("%d.\t%s\n", index+1, sport)
	}
	elapsed := time.Since(start)
	log.Printf("With make took %v", elapsed.Microseconds())

}
func withOutMake() {
	start := time.Now()
	fmt.Println("Start time:", start)
	sports := []string{"Cricket", "Football", "Running", "Badminton"}
	sports = append(sports, "Volleyball")
	sports = append(sports, "Swimming")
	fmt.Println("My favourite sports:")
	for index, sport := range sports {
		fmt.Printf("%d.\t%s\n", index+1, sport)
	}
	sports = append(sports[:3], sports[4:5]...)
	fmt.Println("My most favourite sports:")
	for index, sport := range sports {
		fmt.Printf("%d.\t%s\n", index+1, sport)
	}
	elapsed := time.Since(start)
	log.Printf("Without make took %v", elapsed.Microseconds())
}
