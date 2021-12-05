package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileLineByLine(filePath string) []string {
	lines := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func ConvertStringSliceToInt(lines []string) []int {
	nums := make([]int, 0)
	for _, str := range lines {
		intValue, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, intValue)
	}
	return nums
}
