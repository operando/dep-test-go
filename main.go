package dep_test

import "fmt"

const (
	VERSION = "1.0.0"
)

func Version() {
	fmt.Println(VERSION)
}