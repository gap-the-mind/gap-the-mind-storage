package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gap-the-mind/gap-the-mind-storage/repo"
)

var storeCmd *flag.FlagSet
var queryCmd *flag.FlagSet

var params struct {
	repoPath string
	filePath string
	inMemory bool
}

func init() {

	storeCmd = flag.NewFlagSet("store", flag.ExitOnError)
	storeCmd.StringVar(&params.filePath, "file", "", "path to file to store")
	storeCmd.StringVar(&params.repoPath, "path", "", "path to store")
	storeCmd.BoolVar(&params.inMemory, "memory", false, "in memory store")

	queryCmd = flag.NewFlagSet("query", flag.ExitOnError)
}

func getRepo() (repo.Storage, error) {
	fmt.Printf("%w\n", params)

	if params.inMemory {
		return repo.OpenMemory()
	}

	return repo.OpenFilesystem(params.repoPath)
}

func store(args []string) {
	_, err := getRepo()

	if err != nil {
		log.Fatalf("Unable to open repo %v", err)
	}
}

func query(args []string) {
	_, err := getRepo()

	if err != nil {
		log.Fatalf("Unable to open repo %v", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		log.Fatal("Missing command")
	}

	switch os.Args[1] {

	case storeCmd.Name():
		storeCmd.Parse(os.Args[2:])
		store(storeCmd.Args())

	case queryCmd.Name():
		queryCmd.Parse(os.Args[2:])
		query(queryCmd.Args())

	default:
		flag.PrintDefaults()
		log.Fatalf("Unknown command: %s", os.Args[1])
	}

	flag.Parse()
}
