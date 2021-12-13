package day8

import "strings"

func ParseData(input string) ([][]string, [][]string) {
	inputs := make([][]string, 0)
	outputs := make([][]string, 0)
	split := strings.Split(input, "\n")

	for _, row := range split {
		if row != "" {
			parsed := strings.Split(row, " | ")
			in := strings.Split(parsed[0], " ")
			out := strings.Split(parsed[1], " ")
			inputs = append(inputs, in)
			outputs = append(outputs, out)
		}
	}
	return inputs, outputs
}
