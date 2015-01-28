package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	flag.Parse()

	if err := os.MkdirAll("_vendor", 0755); err != nil {
		log.Fatal(err)
	}

	args := flag.Args()

	env := os.Environ()
	for i, path := range env {
		if strings.HasPrefix(path, "GOPATH=") {
			abs, err := filepath.Abs("_vendor")
			if err != nil {
				log.Fatal(err)
			}
			env[i] = "GOPATH=" + abs
			break
		}
	}

	binary, err := exec.LookPath(args[0])
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Exec(binary, args, env)
	if err != nil {
		log.Fatal(err)
	}
}
