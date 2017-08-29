package utils

import (
	"os"
	"strings"
)

//Dir return the dir part of a path.
//NOTE: not a clean path!
func Dir(path string) string {
	pos := strings.LastIndexByte(path, os.PathSeparator)
	if pos == -1 {
		if os.PathSeparator != '/' {
			pos = strings.LastIndexByte(path, '/')
			if pos == -1 {
				return "."
			}
			return path[:pos+1]
		}
	}

	return path[:pos+1]
}

