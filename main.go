package main

import (
	"flag"
	"fmt"

	day1 "aoc/days/day_1"
	day2 "aoc/days/day_2"
)

func dayRunner(day int) {
	switch day {
	case 1:
		day1.Day1()
	case 2:
		day2.Day2()
	case 3:
		fmt.Println("Day 3")
	case 4:
		fmt.Println("Day 4")
	case 5:
		fmt.Println("Day 5")
	case 6:
		fmt.Println("Day 6")
	case 7:
		fmt.Println("Day 7")
	case 8:
		fmt.Println("Day 8")
	case 9:
		fmt.Println("Day 9")
	case 10:
		fmt.Println("Day 10")
	case 11:
		fmt.Println("Day 11")
	case 12:
		fmt.Println("Day 12")
	case 13:
		fmt.Println("Day 13")
	case 14:
		fmt.Println("Day 14")
	case 15:
		fmt.Println("Day 15")
	case 16:
		fmt.Println("Day 16")
	case 17:
		fmt.Println("Day 17")
	case 18:
		fmt.Println("Day 18")
	case 19:
		fmt.Println("Day 19")
	case 20:
		fmt.Println("Day 20")
	case 21:
		fmt.Println("Day 21")
	case 22:
		fmt.Println("Day 22")
	case 23:
		fmt.Println("Day 23")
	case 24:
		fmt.Println("Day 24")
	case 25:
		fmt.Println("Day 25")
	default:
		fmt.Println("out of luck nerd")
	}
}

func main() {
	var day int
	flag.IntVar(&day, "day", 1, "Day to run")
	flag.Parse()
	dayRunner(day)
}
