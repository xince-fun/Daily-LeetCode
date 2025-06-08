package main

import (
	"bytes"
	"unicode"
)

func clearDigits1(s string) string {
	sbyte := []byte(s)
	for len(sbyte) > 0 {
		idx := bytes.IndexFunc(sbyte, func(r rune) bool {
			return r >= 48 && r <= 57
		})
		if idx == -1 {
			break
		}
		if idx == 0 {
			sbyte = sbyte[1:]
		} else {
			sbyte = append(sbyte[:idx-1], sbyte[idx+1:]...)
		}
	}
	return string(sbyte)
}

func clearDigits(s string) string {
	st := []rune{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
