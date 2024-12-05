package hugo

import "github.com/mmcdole/gofeed"

type Request struct {
	EventType string
	Domain    string
	RssURL    string
	Content   gofeed.Item
}

type Post struct {
	Title      string
	Date       string
	Draft      bool
	Categories []string
	Tags       []string
	Content    string
}
