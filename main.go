package main

import (
	"fmt"
	"os"
)

func die(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func main() {
	die("not implemented")
}
