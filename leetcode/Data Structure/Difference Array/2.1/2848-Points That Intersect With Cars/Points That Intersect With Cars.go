package main

func numberOfPoints(nums [][]int) (ans int) {
	d := [102]int{}
	for _, num := range nums {
		d[num[0]] += 1
		d[num[1]+1] -= 1
	}
	s := 0
	for _, v := range d {
		s += v
		if s > 0 {
			ans++
		}
	}
	return ans
}
