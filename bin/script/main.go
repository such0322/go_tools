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
	switch args[0] {
	case "addMissionDaily":
		addMissionDaily(args[1])
	default:
		fmt.Println("没有对应的方法")
	}

	log.Printf("run time: %s", time.Since(start))
}
