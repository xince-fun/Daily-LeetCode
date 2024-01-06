package main

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]-'a']--
	}
	for _, num := range cnt {
		if num < 0 {
			return false
		}
	}
	return true
}
