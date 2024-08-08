package main

func countBeautifulPairs1(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			num1 := nums[i]
			for num1 >= 10 {
				num1 = num1 % 10
			}
			num2 := nums[j] % 10
			if gcd(num1, num2) == 1 {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countBeautifulPairs(nums []int) (ans int) {
	cnt := [10]int{}
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 {
			x /= 10
		}
		cnt[x]++
	}
	return
}
