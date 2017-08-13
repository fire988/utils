package utils

import (
	"fmt"
	"testing"
)

func TestGenRandomPass(t *testing.T) {
	s := GenRandomPass(16)
	if len(s) != 16 {
		t.Error("length != 16")
	} else {
		fmt.Println("pass:", s)
	}
}

func TestDigiCode(t *testing.T) {
	s := ""
	s = GenRandomDigiCode(6)
	if len(s) != 6 {
		t.Error("error.")
	} else {
		fmt.Println(s)
	}
	s = GenRandomDigiCode(8)
	if len(s) != 8 {
		t.Error("error.")
	} else {
		fmt.Println(s)
	}
	s = GenRandomDigiCode(16)
	if len(s) != 16 {
		t.Error("error.")
	} else {
		fmt.Println(s)
	}
}
