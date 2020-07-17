package main

import (
	"fmt"

	"github.com/gap-the-mind/gap-the-mind-storage/repo"
)

func main() {

	repository, err := repo.Open("../storage")
	err = repository.Reindex()

	fmt.Println(repository, err)

}
