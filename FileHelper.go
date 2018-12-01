package fileHelper


import (
	"os"
	"bufio"
)

func DoOnEachLine(pathToFile string, doer func(line string)) {
	fileHandler, err := os.Open(pathToFile)
	if err != nil {
		panic(err)
	}
	defer fileHandler.Close()

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		go doer(scanner.Text())
	}
}
