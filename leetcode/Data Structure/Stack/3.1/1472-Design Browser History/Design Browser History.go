package main

type BrowserHistory struct {
	st  []string
	st2 []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		st:  []string{homepage},
		st2: []string{},
	}
}

func (t *BrowserHistory) Visit(url string) {
	t.st2 = []string{}
	t.st = append(t.st, url)
}

func (t *BrowserHistory) Back(steps int) (ans string) {
	for i := 0; i < steps && len(t.st) > 1; i++ {
		now := t.st[len(t.st)-1]
		t.st = t.st[:len(t.st)-1]
		t.st2 = append(t.st2, now)
	}
	return t.st[len(t.st)-1]
}

func (t *BrowserHistory) Forward(steps int) (ans string) {
	for i := 0; i < steps && len(t.st2) > 0; i++ {
		now := t.st2[len(t.st2)-1]
		t.st2 = t.st2[:len(t.st2)-1]
		t.st = append(t.st, now)
	}
	return t.st[len(t.st)-1]
}
