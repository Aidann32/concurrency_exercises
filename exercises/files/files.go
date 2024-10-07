package files

import (
	"fmt"
	"github.com/Aidann32/concurrency_exercises/exercises/files/fetcher"
	"github.com/Aidann32/concurrency_exercises/exercises/files/producers_consumer"
	"os"
)

func Run(fileNumber int) {
	if _, err := os.Stat("files/"); os.IsNotExist(err) {
		if err := createFiles(fileNumber); err != nil {
			fmt.Printf("Create directory and files error: %s\n", err)
			return
		}
	}
	_ = os.Chdir("files")
	pr := producers_consumer.NewProducersConsumer(fileNumber, fetcher.FileFetcher{})
	pr.Run()
}
