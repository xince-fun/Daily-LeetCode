package main

// 选或不选的问题

//             dfs(i+1,j-1) + 2 s[i] == s[j]
// dfs(i, j) = max(dfs(i+1,j), dfs(i,j-1)) s[i] != s[j]

func longestPalindrome(s string) string {
	n := len(s)

}
