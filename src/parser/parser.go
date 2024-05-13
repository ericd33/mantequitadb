package parser

import (
	"bufio"
	"fmt"
	"os"
)

func ParseQueryCommand(dbName string) {

	f, error := OpenOrCreateQuery(dbName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == ";" {
			break
		}
		print(line)

	}

}

func OpenOrCreateQuery(fileName string) (*os.File, error) {
	return os.OpenFile(fmt.Sprintf("./%s.q", fileName), os.O_RDONLY, 0644)
}
