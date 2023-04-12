package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	withOutMake()
	withMake()
}
func withMake() {
	start := time.Now()
	fmt.Println("Start time:", start)
	sports := make([]string, 4, 15)
	sports[0] = "Cricket"
	sports[1] = "Football"
	sports[2] = "Running"
	sports[3] = "Badminton"
	sports = append(sports, "Volleyball", "Swimming")
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
	log.Printf("With make took %s", elapsed)

}
func withOutMake() {
	start := time.Now()
	fmt.Println("Start time:", start)
	sports := []string{"Cricket", "Football", "Running", "Badminton"}
	sports = append(sports, "Volleyball", "Swimming")
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
	log.Printf("Without make took %s", elapsed)
}
