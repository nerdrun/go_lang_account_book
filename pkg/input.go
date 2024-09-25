package pkg

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
