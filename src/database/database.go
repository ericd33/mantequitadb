package database

import (
	"fmt"
	"os"
)

type database struct {
	Name string
}

func New(name string) *database {
	return &database{name}
}

func init() {

}

func (db *database) Load(path string) {

	fmt.Printf(path)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (db *database) Create(name string) {
	data := []byte("hellofile")
	err := os.WriteFile("./bin/first.mdb", data, 0644)
	check(err)
}
