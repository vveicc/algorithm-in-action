package main

type StreamChecker struct {
	curr *tire
}

func Constructor(words []string) StreamChecker {
	t := &tire{}
	for _, s := range words {
		t.insert(s)
	}
	t.automate()
	return StreamChecker{t}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.curr = this.curr.children[letter-'a']
	return this.curr.end
}

type tire struct {
	end      bool
	fail     *tire
	children [26]*tire
}

func (t *tire) insert(s string) {
	curr := t
	for _, c := range s {
		c -= 'a'
		if curr.children[c] == nil {
			curr.children[c] = &tire{}
		}
		curr = curr.children[c]
	}
	curr.end = true
}

// 构建 AC 自动机
func (t *tire) automate() {
	t.fail = t
	var q []*tire
	for i, c := range t.children {
		if c == nil {
			t.children[i] = t
		} else {
			c.fail = t
			q = append(q, c)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		o.end = o.end || o.fail.end
		for i, c := range o.children {
			if c == nil {
				o.children[i] = o.fail.children[i]
			} else {
				c.fail = o.fail.children[i]
				q = append(q, c)
			}
		}
	}
}
