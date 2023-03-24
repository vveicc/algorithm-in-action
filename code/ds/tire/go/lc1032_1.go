package main

type StreamChecker struct {
	l int
	t *tire
	s []byte
}

func Constructor(words []string) StreamChecker {
	l := 0
	t := &tire{}
	for _, s := range words {
		t.insert(s)
		l = max(l, len(s))
	}
	return StreamChecker{l, t, nil}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.s = append(this.s, letter-'a')
	if len(this.s) > this.l {
		this.s = this.s[1:]
	}
	return this.t.search(this.s)
}

type tire struct {
	end      bool
	children [26]*tire
}

func (t *tire) insert(s string) {
	curr := t
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i] - 'a'
		if curr.children[c] == nil {
			curr.children[c] = &tire{}
		}
		curr = curr.children[c]
	}
	curr.end = true
}

func (t *tire) search(s []byte) bool {
	curr := t
	for i := len(s) - 1; i >= 0 && curr.children[s[i]] != nil; i-- {
		if curr = curr.children[s[i]]; curr.end {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
