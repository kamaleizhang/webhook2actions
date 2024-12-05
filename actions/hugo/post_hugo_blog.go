package hugo

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

var RepoLocation string

var hugoYamlTemplate = `{{- "---" }}
title: "{{.Title}}"
date: {{.Date}}
draft: false
categories: 
	{{- "[" -}}
	{{- range $index, $id := .Categories -}}
		{{- if ne $index 0 -}},{{- end -}}
		'{{ $id }}'
	{{- end -}}
	{{- "]" }}
tags: {{- "[" -}}
	{{- range $index, $id := .Tags -}}
		{{- if ne $index 0 -}},{{- end -}}
		'{{ $id }}'
	{{- end -}}
	{{- "]" }}
show_comments: true
---

{{.Content}}
`

func HandleHook(w http.ResponseWriter, r *http.Request) {
	//read request body and unmarshal it
	req := &Request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	//generate hugo blog post
	post, err := GenerateHugoBlogPost(req)
	if err != nil {
		http.Error(w, "Error generating hugo blog post", http.StatusInternalServerError)
		return
	}

	//write hugo blog post to file
	year := req.Content.Published[0:4]
	filePath := RepoLocation + "/content/posts/" + year + "/" + req.Content.GUID + ".md"
	err = write2File(filePath, post)
	if err != nil {
		http.Error(w, "Error writing hugo blog post to file", http.StatusInternalServerError)
		return
	}

	//push hugo git repo
	err = pushGitRemote()
	if err != nil {
		http.Error(w, "Error pushing hugo git repo", http.StatusInternalServerError)
		return
	}

	//return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Hugo blog post created"})
}

func GenerateHugoBlogPost(req *Request) (string, error) {
	t, err := template.New("example").Parse(hugoYamlTemplate)
	if err != nil {
		return "", err
	}

	data := Post{
		Title:      req.Content.Title,
		Date:       req.Content.Published,
		Draft:      false,
		Categories: req.Content.Categories,
		Tags:       req.Content.Categories,
		Content:    req.Content.Content,
	}

	var result strings.Builder
	err = t.Execute(&result, data)
	if err != nil {
		return "", err
	}

	finalResult := result.String()
	return finalResult, nil
}

func write2File(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func pushGitRemote() error {
	//git push hugo git repo
	repo, err := git.PlainOpen(RepoLocation)
	if err != nil {
		return err
	}
	wt, err := repo.Worktree()
	if err != nil {
		return err
	}
	_, err = wt.Add(RepoLocation)
	if err != nil {
		return err
	}
	_, err = wt.Commit("Add new blog post", &git.CommitOptions{})
	if err != nil {
		return err
	}
	err = repo.Push(&git.PushOptions{})
	if err != nil {
		return err
	}
	return nil
}
