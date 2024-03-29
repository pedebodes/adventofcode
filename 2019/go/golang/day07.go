package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type OpType int

const (
	OpAdd OpType = iota + 1
	OpMul
	OpInput
	OpOutput
	OpJumpTrue
	OpJumpFalse
	OpLessThan
	OpEquals
	OpEnd = 99
)

type ModeType int

const (
	ModePosition  ModeType = 0
	ModeImmediate          = 1
)

func getArgPos(ops []int, pos int, mode ModeType) (int, error) {
	if pos >= len(ops) {
		return -1, fmt.Errorf("invalido: %d", pos)
	}
	if mode == ModePosition {
		if pos >= len(ops) {
			return -1, fmt.Errorf("invalido: %d", pos)
		}
		return ops[pos], nil
	}
	return pos, nil
}

type Amplifier struct {
	ops    []int
	pc     int
	Halted bool
}

func (amp *Amplifier) Iterate(input func() int, yieldOnInput bool, debug bool) (int, bool, error) {
	log := ioutil.Discard
	if debug {
		log = os.Stdout
	}
	for amp.pc < len(amp.ops) {
		op := OpType(amp.ops[amp.pc])
		if op == OpEnd {
			fmt.Fprintf(log, "parou!\n")
			amp.Halted = true
			return 0, true, nil
		}

		aMode := ModePosition
		bMode := ModePosition
		if op > 99 {
			aMode = ModeType((op / 100) % 10)
			bMode = ModeType((op / 1000) % 10)
			op %= 100
		}

		switch op {
		case OpAdd:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := getArgPos(amp.ops, amp.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := getArgPos(amp.ops, amp.pc+3, ModePosition)
			if err != nil {
				return 0, false, err
			}
			amp.ops[o] = amp.ops[a] + amp.ops[b]
			fmt.Fprintf(log, "set %d @%d\n", amp.ops[o], o)
			amp.pc += 4
		case OpMul:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := getArgPos(amp.ops, amp.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := getArgPos(amp.ops, amp.pc+3, ModePosition)
			if err != nil {
				return 0, false, err
			}
			amp.ops[o] = amp.ops[a] * amp.ops[b]
			fmt.Fprintf(log, "set %d @%d\n", amp.ops[o], o)
			amp.pc += 4
		case OpInput:
			a, err := getArgPos(amp.ops, amp.pc+1, ModePosition)
			if err != nil {
				return 0, false, err
			}
			amp.ops[a] = input()
			fmt.Fprintf(log, "entrada: %d @%d\n", amp.ops[a], a)
			amp.pc += 2
			if yieldOnInput {
				return 0, false, nil
			}
		case OpOutput:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			fmt.Fprintf(log, "Saida: %d @%d\n", amp.ops[a], a)
			amp.pc += 2
			return amp.ops[a], false, nil
		case OpJumpTrue:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := getArgPos(amp.ops, amp.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			if amp.ops[a] != 0 {
				amp.pc = amp.ops[b]
				fmt.Fprintf(log, "-> %d\n", amp.pc)
			} else {
				amp.pc += 3
			}
		case OpJumpFalse:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := getArgPos(amp.ops, amp.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			if amp.ops[a] == 0 {
				amp.pc = amp.ops[b]
				fmt.Fprintf(log, "-> %d\n", amp.pc)
			} else {
				amp.pc += 3
			}
		case OpLessThan:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := getArgPos(amp.ops, amp.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := getArgPos(amp.ops, amp.pc+3, ModePosition)
			if err != nil {
				return 0, false, err
			}
			if amp.ops[a] < amp.ops[b] {
				amp.ops[o] = 1
			} else {
				amp.ops[o] = 0
			}
			amp.pc += 4
		case OpEquals:
			a, err := getArgPos(amp.ops, amp.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := getArgPos(amp.ops, amp.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := getArgPos(amp.ops, amp.pc+3, ModePosition)
			if err != nil {
				return 0, false, err
			}
			if amp.ops[a] == amp.ops[b] {
				amp.ops[o] = 1
			} else {
				amp.ops[o] = 0
			}
			amp.pc += 4
		default:
			return 0, false, fmt.Errorf("invalido: %d", op)
		}
	}
	return 0, false, nil
}

func RunCircuit(ops []int, inputs []int, debug bool) (int, error) {
	amps := make([]Amplifier, len(inputs))
	for i := range amps {
		amps[i].ops = make([]int, len(ops))
		copy(amps[i].ops, ops)
	}

	for i := range amps {
		_, _, err := amps[i].Iterate(func() int {
			return inputs[i]
		}, true, debug)
		if err != nil {
			return 0, fmt.Errorf("erro execucao %w", err)
		}
	}

	var output int
	var halt bool
	var count int
	for !halt {
		for i := range amps {
			var o int
			var err error
			o, halt, err = amps[i].Iterate(func() int {
				return output
			}, false, debug)
			if err != nil {
				return 0, fmt.Errorf("erro execucao %w", err)
			}
			if halt {
				break
			}
			output = o
		}

		if debug {
			count++
			fmt.Printf("interacao: %d -> %d\n", count, output)
		}
	}

	return output, nil
}

func verifyInput(input []int) bool {
	seen := make(map[int]bool)
	for _, i := range input {
		if _, ok := seen[i]; ok {
			return false
		}
		seen[i] = true
	}
	return true
}

func FindPhaseSettings(ops []int, lo, hi int) (int, error) {
	result := 0

	for i := lo; i <= hi; i++ {
		for j := lo; j <= hi; j++ {
			for k := lo; k <= hi; k++ {
				for l := lo; l <= hi; l++ {
					for m := lo; m <= hi; m++ {
						input := []int{i, j, k, l, m}
						if !verifyInput(input) {
							continue
						}

						amp, err := RunCircuit(ops, input, false)
						if err != nil {
							return 0, fmt.Errorf("Erro execucao: %w", err)
						}
						if amp > result {
							result = amp
						}
					}
				}
			}
		}
	}

	return result, nil
}

func main() {
	input, err := os.Open("inputs/input07.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("erro abrir arquivo: %w", err))
	}
	defer input.Close()

	raw, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(fmt.Errorf("erro leitura arquivo: %w", err))
	}
	ops, err := func(val []byte) ([]int, error) {
		vals := strings.Split(strings.TrimSpace(string(val)), ",")
		result := make([]int, len(vals))
		for i := range vals {
			result[i], err = strconv.Atoi(vals[i])
			if err != nil {
				return nil, fmt.Errorf("erro ao converter: %w", err)
			}
		}
		return result, nil
	}(raw)
	if err != nil {
		log.Fatal(fmt.Errorf("Falha: %w", err))
	}

	

	// Parte 1
	result, err := FindPhaseSettings(ops, 0, 4)
	if err != nil {
		log.Fatal(fmt.Errorf("Erro parte 1: %w", err))
	}
	fmt.Printf("Parte 1: %d\n", result)

	// Parte 2
	result, err = FindPhaseSettings(ops, 5, 9)
	if err != nil {
		log.Fatal(fmt.Errorf("Erro parte 2: %w", err))
	}
	fmt.Printf("Parte 2: %d\n", result)
}