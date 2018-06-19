package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func cxx(file string) result {
	var total, empty, effective, comment int64
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gloc: cxx: %v\n", err)
	}
	// per char
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanRunes)

	// comment related
	var inSingleComment, inMultiComment, onceComment bool

	// string, valid source line
	var inString, validSrc bool
	var last, cur byte
	cur = '\n'
	for input.Scan() {
		c := input.Bytes()
		if len(c) != 1 {
			continue
		}
		last, cur = cur, c[0]
		if c[0] == '\n' {
			total++
			if last == '\n' {
				empty++
			}
			// current line is valid?
			if validSrc && last != '\\' {
				//fmt.Println(total)
				effective++
			}
			if onceComment || inMultiComment || inSingleComment {
				comment++
			}
			if !inString {
				validSrc = false
			}
			onceComment = false
			inSingleComment = false
			continue
		}
		// skip all chars in a string
		if inString && cur != '"' {
			continue
		}
		switch cur {
		case '"':
			// start or end of a string
			inString = !inString
		case '*':
			if last == '/' {
				inMultiComment = true
			}
		case '/':
			if last == '*' {
				inMultiComment = false
				onceComment = true
			} else if last == '/' {
				inSingleComment = true
			}
		case '\t', '\v', '\f', '\r', ' ':
			// skip white spaces
		default:
			// current line is valid?
			if !inSingleComment && !inMultiComment {
				validSrc = true
			}
		}
	}

	// generate relative path
	rel, _ := filepath.Rel(project, file)
	return result{rel, total, empty, effective, comment}
}
