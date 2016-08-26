package main

import (
	"flag"
	"log"

	"github.com/umaumax/goafplay"
)

var ()

func init() {
}

func main() {
	flag.Parse()

	for _, v := range flag.Args() {
		err := goafplay.Do(v)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
