// +build !go1.1

package main

import "fmt"

func init() {
	fmt.Println("go < 1.1 doesn't have any version tags")
}
