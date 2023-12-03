package day1

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

// i love go

func Day1() {
	start := time.Now()
	input, err := os.Open("days/day_1/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	demDigits := 0

	for s.Scan() {
		first, last := 0, 0

		for _, c := range s.Text() {
			if unicode.IsDigit(c) {
				// old school c style char to int conversion
				last = int(c - '0')
				if first == 0 {
					first = last
				}
			}
		}

		d := fmt.Sprintf("%d%d", first, last)

		dInt := 0
		fmt.Sscanf(d, "%d", &dInt)
		if err != nil {
			panic(err)
		}

		demDigits = demDigits + dInt

	}

	fmt.Println(demDigits)

	elapse := time.Since(start)
	fmt.Printf("took %s\n", elapse)

}
