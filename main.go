package main

import (
	"flag"
	"fmt"
	"go-testing/words"
	"log"
	"os"
)

var (
	file   string
	filter string
)

func init() {
	flag.StringVar(&file, "file", "", "arquivo com o texto")
	flag.StringVar(&filter, "filter", "", "filtra pela palavra")
	flag.Parse()
}

func main() {
	if file != "" {
		panic("missing file")
	}
	content, err := os.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	c := words.Count(string(content), filter)

	fmt.Println("total:", c)
}
