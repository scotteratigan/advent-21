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
	switch day {
	case "1":
		nums := ConvertStringSliceToInt(ReadFileLineByLine("./input01"))
		switch problem {
		case "1":
			D1P1(nums)
		case "2":
			D1P2(nums)
		default:
			logInvalidProblemNumber(problem)
		}
	case "2":
		maneuvers := readManeuvers(ReadFileLineByLine("./input02"))
		switch problem {
		case "1":
			D2P1(maneuvers)
		case "2":
			logInvalidProblemNumber(problem)
		default:
			logInvalidProblemNumber(problem)
		}
	default:
		log.Fatal("Day " + day + " not coded yet!")
	}
}

func logInvalidProblemNumber(n string) {
	log.Fatal("Problem number " + n + " is invalid! Valid values are 1-2.")
}
