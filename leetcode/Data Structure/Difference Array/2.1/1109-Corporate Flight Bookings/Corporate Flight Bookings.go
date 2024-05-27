package main

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := []int{}
	d := make([]int, n+2)

	for _, b := range bookings {
		d[b[0]] += b[2]
		d[b[1]+1] -= b[2]
	}
	s := 0
	for i := 1; i <= n; i++ {
		s += d[i]
		ans = append(ans, s)
	}
	return ans
}
