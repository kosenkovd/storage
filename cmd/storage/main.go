package main

import (
	"fmt"
	"log"

	"github.com/kosenkovd/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("testest"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("uploaded", file)

	fileFromCache, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("restored", fileFromCache)
}
