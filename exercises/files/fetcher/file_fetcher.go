package fetcher

import (
	"os"
	"strconv"
	"strings"
)

type FileFetcher struct {
}

func (f *FileFetcher) Fetch(resource string) (output FetchOutput) {
	data, _ := os.ReadFile(resource)
	input := strings.Split(string(data), "\n")
	output.Operation = input[0]
	operands := strings.Split(input[1], " ")
	output.Operand1, _ = strconv.ParseFloat(operands[0], 64)
	output.Operand2, _ = strconv.ParseFloat(operands[1], 64)
	return
}
