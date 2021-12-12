package main

import (
	"fmt"
	"log"
	"math"
)

/*

--- Day 3: Binary Diagnostic ---
The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

So, the gamma rate is the binary number 10110, or 22 in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.

Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine? (Be sure to represent your answer in decimal, not binary.)

*/

const (
	zero = 48
	one  = 49
)

type count struct {
	zeros int
	ones  int
}

func D3P1(inputs []string) {
	counts := GetDigitCounts(inputs)
	gammaRate := BinStrToInt(strOfMostCommonValues(counts))
	epsilonRate := BinStrToInt(strOfLeastCommonValues(counts))
	powerConsumption := gammaRate * epsilonRate
	fmt.Println("Power consumption:", powerConsumption)

}

func GetDigitCounts(inputs []string) []count {
	// reports count of zeros vs ones in binary strings
	// 111 100 101 -> {3 0} {1 2} {2 1}
	inputLen := len(inputs[0])
	counts := make([]count, 0)
	for i := 0; i < inputLen; i++ {
		c := count{}
		for _, input := range inputs {
			switch input[i] {
			case zero:
				c.zeros++
			case one:
				c.ones++
			default:
				log.Fatal("Unknown digit encountered: " + string(input[0]))
			}
		}
		counts = append(counts, c)
	}
	return counts
}

func strOfMostCommonValues(counts []count) string {
	var result string
	for _, c := range counts {
		if c.ones > c.zeros {
			result += "1"
		} else if c.zeros > c.ones {
			result += "0"
		} else {
			log.Fatal("equal number of zeroes and ones!")
		}
	}
	return result
}

func strOfLeastCommonValues(counts []count) string {
	var result string
	for _, c := range counts {
		if c.ones < c.zeros {
			result += "1"
		} else if c.zeros < c.ones {
			result += "0"
		} else {
			log.Fatal("equal number of zeroes and ones!")
		}
	}
	return result
}

func BinStrToInt(input string) int {
	total := 0
	var power float64 = -1
	for i := len(input) - 1; i >= 0; i-- {
		power += 1.0
		v := input[i]
		switch v {
		case zero:
		case one:
			total += int(math.Pow(2, power))
		default:
			log.Fatal("Invalid value specified: " + string(v))
		}
	}
	return total
}
