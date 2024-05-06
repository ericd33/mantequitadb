package main

import (
	D "gotest/src/database"
)

func main() {
	db := D.New("Roberto")
	db.Create("first.mdb")
	// db.Load("hola")
}
