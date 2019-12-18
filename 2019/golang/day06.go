package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readLines("inputs/input6.txt")

	orbits := make(map[string]string)
	for _, value := range input {
		parts := strings.Split(value, ")")
		orbits[parts[1]] = parts[0]
	}
	{
		
		output := doPart1(orbits)
		fmt.Println("Parte 1: ",output)
	}
	{
		output := doPart2(orbits)
		fmt.Println("Parte 2: ",output)
	}
}

func doPart1(orbits map[string]string) int {
	total := 0
	for object := range orbits {
		for {
			parent, ok := orbits[object]
			if !ok {
				break
			}
			object = parent
			total++
		}
	}

	return total
}

func doPart2(orbits map[string]string) int {
	path := make(map[string]int)

	object, distance := orbits["YOU"], 0
	for {
		path[object] = distance
		parent, ok := orbits[object]
		if !ok {
			break
		}
		object = parent
		distance++
	}

	object, distance = orbits["SAN"], 0
	for {
		pathDistance, ok := path[object]
		if ok {
			distance += pathDistance
			break
		}
		parent, ok := orbits[object]
		if !ok {
			panic("YOU and SAN nao conectado")
		}
		object = parent
		distance++
	}

	return distance
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}