package main

import "container/heap"

type Twitter struct {
	userMap map[int]*User
}

type Tweet struct {
	Id, Time int
	next     *Tweet
}

func Constructor() Twitter {
	return Twitter{make(map[int]*User)}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	if u, ok := t.userMap[userId]; ok {
		u.post(tweetId)
	} else {
		u = NewUser(userId)
		t.userMap[userId] = u

		u.post(tweetId)
	}
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	tweetIds := make([]int, 0)

	if u, ok := t.userMap[userId]; ok {
		pq := make(PriorityQueue, 0)
		heap.Init(&pq)

		for followee := range u.Following {
			headT := t.userMap[followee].Head
			if headT != nil {
				heap.Push(&pq, headT)
			}
		}

		for pq.Len() > 0 {
			if len(tweetIds) == 10 {
				break
			}
			tweet := heap.Pop(&pq).(*Tweet)
			if tweet.next != nil {
				heap.Push(&pq, tweet.next)
			}

			tweetIds = append(tweetIds, tweet.Id)
		}
	}
	return tweetIds
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.userMap[followerId]; !ok {
		f := NewUser(followerId)
		t.userMap[followerId] = f
	}
	if _, ok := t.userMap[followeeId]; !ok {
		f := NewUser(followeeId)
		t.userMap[followeeId] = f
	}
	t.userMap[followerId].follow(followeeId)
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if f, ok := t.userMap[followerId]; ok {
		f.unfollow(followeeId)
	}
}

var timestamp int = 0

type User struct {
	Id        int
	Following map[int]bool
	Head      *Tweet
}

func NewUser(userId int) *User {
	u := &User{
		Id:        userId,
		Following: make(map[int]bool),
	}
	u.follow(userId) // 关注自己
	return u
}

func (u *User) follow(userId int) {
	u.Following[userId] = true
}

func (u *User) unfollow(userId int) {
	if userId != u.Id {
		delete(u.Following, userId)
	}
}

func (u *User) post(tweetId int) {
	tweet := &Tweet{tweetId, timestamp, nil}
	timestamp++
	tweet.next = u.Head
	u.Head = tweet
}

type PriorityQueue []*Tweet

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 按时间降序排列
	return pq[i].Time > pq[j].Time
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	tweet := x.(*Tweet)
	*pq = append(*pq, tweet)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	tweet := old[n-1]
	old[n-1] = nil // 防止内存泄漏
	*pq = old[:n-1]
	return tweet
}
