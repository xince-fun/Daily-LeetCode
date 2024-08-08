package main

import "sort"

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	n := len(watchedVideos)
	vis := make([]bool, n)

	g := make([][]int, n)
	for i, v := range friends {
		g[i] = append(g[i], v...)
	}

	cnt := 0
	q := []int{id}
	vis[id] = true
	for len(q) > 0 {
		cnt++
		tmp := q
		q = []int{}
		for _, p := range tmp {
			for _, f := range g[p] {
				if !vis[f] {
					vis[f] = true
					q = append(q, f)
				}
			}
		}
		if cnt == level {
			break
		}
	}

	m := map[string]int{}
	for _, p := range q {
		for _, video := range watchedVideos[p] {
			m[video]++
		}
	}
	res := []string{}
	for video := range m {
		res = append(res, video)
	}
	sort.Slice(res, func(i, j int) bool {
		return m[res[i]] < m[res[j]] || m[res[i]] == m[res[j]] && res[i] < res[j]
	})

	return res
}
