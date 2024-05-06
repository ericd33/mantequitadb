package main

import (
	D "mantequitadb/src/database"
)

var Debug bool = true

func main() {
	db := D.New("Roberto")
	db.Create("first")
	// db.Load("hola")
}
