package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/eatonphil/riscve"
)

func main() {
	pf := flag.String("program", "", "program to run")
	flag.Parse()

	if pf == nil || *pf == "" {
		flag.Usage()
		return
	}

	f, err := ioutil.ReadFile(*pf)
	if err != nil {
		log.Fatalf("Failed to read program: %s", err.Error())
	}

	p, err := riscve.Parse(f)
	if err != nil {
		log.Fatalf("Failed to parse program: %s", err.Error())
	}

	err = riscve.Emulate(p)
	if err != nil {
		log.Fatalf("Failed emulating program: %s", err.Error())
	}
}
