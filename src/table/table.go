package table

import "fmt"

type table struct {
	name string
}

func New(name string) *table {
	return &table{name}
}

func init() {

}

func (db *table) Load(path string) {

	fmt.Printf(path)
}
