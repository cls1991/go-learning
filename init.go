package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == "" {
		log.Fatal("$USER is not set")

	}
	if home == "" {
		home = "/home/" + user

	}
	if gopath == "" {
		gopath = home + "/go"

	}
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")

}

func main() {
	fmt.Println("constants are:", KB, MB, GB, TB, PB)
	fmt.Println("vars are:", home, user, gopath)

}
