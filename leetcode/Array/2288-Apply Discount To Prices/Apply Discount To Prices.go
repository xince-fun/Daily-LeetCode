package main

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	strs := strings.Split(sentence, " ")
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		if str[0] == '$' && len(str) > 1 {
			num, err := strconv.Atoi(str[1:])
			if err == nil {
				f := float64(num) * float64((100 - discount)) / 100.0
				strs[i] = fmt.Sprintf("$%.2f", f)
			}
		}
	}
	return strings.Join(strs, " ")
}
