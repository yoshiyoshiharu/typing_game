package typing

import (
	"os"
	"bufio"
	"fmt"
)

func Judge() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		in := scanner.Text()
	  fmt.Println("in: ", in)
	}
}
