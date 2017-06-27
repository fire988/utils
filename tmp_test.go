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
