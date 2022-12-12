package main

import (
	"fmt"
	"testing"
)

func TestGetInterfaces(t *testing.T) {
	addr, err := GetIfeces()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v+", addr)
}
