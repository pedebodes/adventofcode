package main

import "fmt"

func main() {
	input := []int{3,8,1005,8,310,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,1001,8,0,29,1,2,11,10,1,1101,2,10,2,1008,18,10,2,106,3,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,102,1,8,67,2,105,15,10,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,1001,8,0,93,2,1001,16,10,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,102,1,8,119,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,141,2,7,17,10,1,1103,16,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,102,1,8,170,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1002,8,1,193,1,7,15,10,2,105,13,10,1006,0,92,1006,0,99,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,101,0,8,228,1,3,11,10,1006,0,14,1006,0,71,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,101,0,8,261,2,2,2,10,1006,0,4,3,8,102,-1,8,10,101,1,10,10,4,10,108,0,8,10,4,10,101,0,8,289,101,1,9,9,1007,9,1049,10,1005,10,15,99,109,632,104,0,104,1,21101,0,387240009756,1,21101,327,0,0,1105,1,431,21101,0,387239486208,1,21102,1,338,0,1106,0,431,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21102,3224472579,1,1,21101,0,385,0,1106,0,431,21101,0,206253952003,1,21102,396,1,0,1105,1,431,3,10,104,0,104,0,3,10,104,0,104,0,21102,709052072296,1,1,21102,419,1,0,1105,1,431,21102,1,709051962212,1,21102,430,1,0,1106,0,431,99,109,2,21202,-1,1,1,21102,1,40,2,21102,462,1,3,21102,452,1,0,1105,1,495,109,-2,2105,1,0,0,1,0,0,1,109,2,3,10,204,-1,1001,457,458,473,4,0,1001,457,1,457,108,4,457,10,1006,10,489,1101,0,0,457,109,-2,2105,1,0,0,109,4,2102,1,-1,494,1207,-3,0,10,1006,10,512,21101,0,0,-3,22101,0,-3,1,21202,-2,1,2,21102,1,1,3,21101,531,0,0,1105,1,536,109,-4,2106,0,0,109,5,1207,-3,1,10,1006,10,559,2207,-4,-2,10,1006,10,559,21202,-4,1,-4,1105,1,627,22102,1,-4,1,21201,-3,-1,2,21202,-2,2,3,21102,1,578,0,1105,1,536,21202,1,1,-4,21102,1,1,-1,2207,-4,-2,10,1006,10,597,21101,0,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,619,21201,-1,0,1,21102,1,619,0,106,0,494,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2106,0,0}

	r := robot(0)
	r.runIntcode(input)
	fmt.Println("Part 1 : ", len(r.paintMap))

	fmt.Println("-------- PART 2 --------")
	r = robot(1)
	r.runIntcode(input)
	printMap(r.paintMap)
}

func (r Robot) runIntcode(inputArr []int) {
	input := *getInputMap(inputArr)
	i := 0
	relativeOffset := 0
	colorToPaint := -1
	for true {
		instruction := input[i]

		opcode := instruction % 100
		instruction /= 100

		if opcode == 99 {
			return
		} else if opcode == 3 {
			targetIndex := input[i+1]
			if instruction%10 == 2 {
				targetIndex += relativeOffset
			}

			input[targetIndex] = r.input

			i += 2
		} else if opcode == 4 {
			index := input[i+1]
			val := index
			if instruction%10 == 0 {
				val = input[index]
			} else if instruction%10 == 2 {
				val = input[index+relativeOffset]
			}
			// fmt.Print(val)
			if colorToPaint != -1 {
				r.paint(colorToPaint)
				r.turn(val)
				colorToPaint = -1
			} else {
				colorToPaint = val
			}

			i += 2
		} else if opcode == 9 {
			index := input[i+1]
			mode := instruction % 10
			val := index
			if mode == 0 {
				val = input[index]
			} else if mode == 2 {
				val = input[val+relativeOffset]
			}
			relativeOffset += val
			i += 2
		} else {
			param1 := input[i+1]
			mode1 := instruction % 10

			instruction /= 10

			param2 := input[i+2]
			mode2 := instruction % 10

			instruction /= 10

			targetIndex := input[i+3]
			mode3 := instruction % 10

			if mode1 == 0 {
				param1 = input[param1]
			} else if mode1 == 2 {
				param1 = input[param1+relativeOffset]
			}

			if mode2 == 0 {
				param2 = input[param2]
			} else if mode2 == 2 {
				param2 = input[param2+relativeOffset]
			}

			if mode3 == 2 {
				targetIndex += relativeOffset
			}

			if opcode == 1 {
				input[targetIndex] = param1 + param2
				i += 4
			} else if opcode == 2 {
				input[targetIndex] = param1 * param2
				i += 4
			} else if opcode == 5 {
				if param1 != 0 {
					i = param2
				} else {
					i += 3
				}
			} else if opcode == 6 {
				if param1 == 0 {
					i = param2
				} else {
					i += 3
				}
			} else if opcode == 7 {
				if param1 < param2 {
					input[targetIndex] = 1
				} else {
					input[targetIndex] = 0
				}
				i += 4
			} else if opcode == 8 {
				if param1 == param2 {
					input[targetIndex] = 1
				} else {
					input[targetIndex] = 0
				}
				i += 4
			}
		}
	}
}

func getInputMap(arr []int) *map[int]int {
	m := make(map[int]int)
	for i, val := range arr {
		m[i] = val
	}
	return &m
}

func (r *Robot) paint(color int) {
	r.paintMap[r.currentLocation] = color
	r.coordinatesSeen[r.currentLocation] = true
}

func (r *Robot) turn(direction int) {
	switch r.currentDirection {
	case "N":
		if direction == 0 {
			r.currentLocation.col = r.currentLocation.col - 1
			r.currentDirection = "W"
		} else {
			r.currentLocation.col = r.currentLocation.col + 1
			r.currentDirection = "E"
		}
	case "W":
		if direction == 0 {
			r.currentLocation.row = r.currentLocation.row - 1
			r.currentDirection = "S"
		} else {
			r.currentLocation.row = r.currentLocation.row + 1
			r.currentDirection = "N"
		}
	case "S":
		if direction == 0 {
			r.currentLocation.col = r.currentLocation.col + 1
			r.currentDirection = "E"
		} else {
			r.currentLocation.col = r.currentLocation.col - 1
			r.currentDirection = "W"
		}
	case "E":
		if direction == 0 {
			r.currentLocation.row = r.currentLocation.row + 1
			r.currentDirection = "N"
		} else {
			r.currentLocation.row = r.currentLocation.row - 1
			r.currentDirection = "S"
		}
	}
	r.input = r.paintMap[r.currentLocation]
}

func printMap(imgMap map[Coordinates]int) {
	maxRow := 0
	maxCol := 0
	minRow := 0
	minCol := 0
	for coords := range imgMap {
		if coords.row > maxRow {
			maxRow = coords.row
		}
		if coords.row < minRow {
			minRow = coords.row
		}
		if coords.col > maxCol {
			maxCol = coords.col
		}
		if coords.col < minCol {
			minCol = coords.col
		}
	}

	for row := maxRow; row >= minRow; row-- {
		for col := minCol; col <= maxCol; col++ {
			if imgMap[Coordinates{row, col}] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func robot(startingColor int) *Robot {
	return &Robot{
		Coordinates{0, 0},
		"N",
		map[Coordinates]int{Coordinates{0, 0}: startingColor},
		map[Coordinates]bool{},
		startingColor,
	}
}

type Robot struct {
	currentLocation  Coordinates
	currentDirection string
	paintMap         map[Coordinates]int
	coordinatesSeen  map[Coordinates]bool
	input            int
}

type Coordinates struct {
	row int
	col int
}