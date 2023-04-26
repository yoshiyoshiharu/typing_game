package typing

import (
	"os"
	"bufio"
)

func Judge(title string) bool {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()

	if title == input {
		return true
	} else {
		return false
	}
}
