package main

import "strings"

func simplifyPath1(path string) (ans string) {
	st := []string{}
	low, fast := 0, 0
	for fast < len(path) {
		if path[fast] == '/' {
			sub := path[low:fast]
			low = fast + 1
			if sub == ".." && len(st) > 0 {
				st = st[:len(st)-1]
			} else if sub == "." {
			} else if sub != "" {
				st = append(st, sub)
			}
		}
		fast++
	}
	for i := 0; i < len(st); i++ {
		ans += "/"
		ans += st[i]
	}
	return
}

func simplifyPath(path string) (ans string) {
	paths := strings.Split(path, "/")

	st := []string{}
	for _, p := range paths {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
			continue
		}
		st = append(st, p)
	}
	return "/" + strings.Join(st, "/")
}
