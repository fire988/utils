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

func TestPersonID(t *testing.T) {
	if IsValidPersonID("330107216911255451") {
		t.Error("error checking 330107216911255451")
	}
	if IsValidPersonID("290121186901020824") {
		t.Error("error checking 290121186901020824")
	}
}

func TestStringDisorder(t *testing.T) {
	r := StringDisorder("0123456789")
	if len(r) != len("0123456789") {
		t.Error("error 01")
	}
	t.Log(r)
}
