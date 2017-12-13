package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]
	fmt.Println(args)
	switch args[0] {
	case "addMissionDaily":
		addMissionDaily(args[1])

	}
	log.Printf("run time: %s", time.Since(start))
}
