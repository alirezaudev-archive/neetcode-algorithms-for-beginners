package main

// https://leetcode.com/problems/design-browser-history/

type Page struct {
	url  string
	next *Page
	prev *Page
}

type BrowserHistory struct {
	current *Page
}

func constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		current: &Page{url: homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	newPage := &Page{url: url}
	newPage.prev = this.current
	this.current.next = newPage
	this.current = newPage
}

func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && this.current.prev != nil {
		this.current = this.current.prev
		steps--
	}

	return this.current.url
}

func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && this.current.next != nil {
		this.current = this.current.next
		steps--
	}

	return this.current.url
}
