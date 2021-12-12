package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("Please specify the day (1) and the problem (1-2) to execute.")
	}
	day := args[0]
	problem := args[1]
	if problem != "1" && problem != "2" {
		log.Fatal("Problem number " + problem + " is invalid! Valid values are 1-2.")
	}
	switch day {
	case "1":
		nums := ConvertStringSliceToInt(ReadFileLineByLine("./input01"))
		switch problem {
		case "1":
			D1P1(nums)
		case "2":
			D1P2(nums)
		}
	case "2":
		maneuvers := readManeuvers(ReadFileLineByLine("./input02"))
		switch problem {
		case "1":
			D2P1(maneuvers)
		case "2":
			D2P2(maneuvers)
		}
	case "3":
		instructions := ReadFileLineByLine("./input03")
		switch problem {
		case "1":
			D3P1(instructions)
		case "2":
		}
	default:
		log.Fatal("Day " + day + " not coded yet!")
	}
}
