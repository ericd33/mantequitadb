package database

import (
	"fmt"
	"mantequitadb/src/config"
	"os"
)

type database struct {
	Name         string
	DatabaseFile *os.File
}

func New(name string) *database {
	return &database{name, nil}
}

// createBinFolderIfNotExists manages the creation of the bin folder needed for the database
func createBinFolderIfNotExists() {
	_, err := os.Stat("./bin")
	if err != nil {
		if err.Error() == "CreateFile ./bin: The system cannot find the file specified." {
			config.DebugIt("Bin folder does not exist, Creating...")
			os.Mkdir("bin", 0644)
			return
		}
		return
	}

	config.DebugIt("Bin folder already exists, continue...")
	return
}

func openOrCreateDbFile(name string) *os.File {
	f, err := os.OpenFile(fmt.Sprintf("./bin/%s.mdb", name), os.O_CREATE, 0644)
	if err != nil {
		config.DebugIt("Something's wrong")
	}
	return f

}

// Load loads the database into the memory
func (db *database) Load(path string) {
	fmt.Printf(path)
}

// check throws an error in the console if there was one
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Create creates a new database file with the given name
func (db *database) Create(name string) {
	createBinFolderIfNotExists()
	openOrCreateDbFile(name)
}

// AppendByteData appends data in bytes to the file
func (db *database) AppendByteData(data []byte) {
	df := db.DatabaseFile

	fi, err := df.Stat()
	if err != nil {
		panic(err)
	}

	df.WriteAt(data, fi.Size())
}

func (db *database) AppendString(data string) {
	bd := []byte(data)
	db.AppendByteData(bd)
}
