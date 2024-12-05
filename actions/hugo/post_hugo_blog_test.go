package hugo

import (
	"testing"

	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
)

func TestGenerateHugoBlogPost(t *testing.T) {
	req := Request{
		EventType: "test",
		Domain:    "test.com",
		RssURL:    "test.com",
		Content: gofeed.Item{
			Title:      "test tile",
			Content:    "test content",
			Categories: []string{"test2", "test1"},
			Published:  "2023-04-25T14:30:00Z",
		},
	}
	result, err := GenerateHugoBlogPost(req)
	if err != nil {
		t.Errorf("Error generating Hugo blog post: %v", err)
	}
	expect := `---
title: "test tile"
date: 2023-04-25T14:30:00Z
draft: false
categories:['test2','test1']
tags:['test2','test1']
show_comments: true
---

test content
`
	assert.Equal(t, expect, result)
}
