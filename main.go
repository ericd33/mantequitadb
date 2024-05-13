package main

import (
	"flag"
	// "fmt"
	"mantequitadb/src/parser"
	// D "mantequitadb/src/database"
)

var Debug bool = true

func main() {
	cmdQuery := flag.String("q", "", "")
	flag.Parse()

	if cmdQuery != nil {
		parser.ParseQueryCommand(string(*cmdQuery))
	}

	// fmt.Printf("my cmd: \"%v\"\n", string(*cmd))
	// db := D.New("Roberto")
	// db.Create("first")
	// db.Load("hola")
}
