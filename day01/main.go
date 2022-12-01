package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() error {
	input, err := os.ReadFile("day01/input.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")

	total_calories := 0
	fattest_elf_calories := 0
	for _, line := range lines {
		if line == "" {
			if total_calories > fattest_elf_calories {
				fattest_elf_calories = total_calories
			}
			total_calories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		total_calories += calories
	}

	fmt.Printf("%d", fattest_elf_calories)
	return nil
}

func Part2() error {
	input, err := os.ReadFile("day01/input.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")

	total_calories := 0
	fattest_elf_calories := make([]int, 3)
	for _, line := range lines {
		if line == "" {
			for i, c := range fattest_elf_calories {
				if total_calories > c {
					for j := len(fattest_elf_calories) - 1; j > i; j-- {
						fattest_elf_calories[j] = fattest_elf_calories[j-1]
					}
					fattest_elf_calories[i] = total_calories
					break
				}
			}
			fmt.Printf("%v\n", fattest_elf_calories)
			total_calories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		total_calories += calories
	}

	fmt.Printf("%d", fattest_elf_calories[0]+fattest_elf_calories[1]+fattest_elf_calories[2])
	return nil
}
