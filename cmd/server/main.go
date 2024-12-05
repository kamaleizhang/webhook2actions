package main

import (
	"flag"
	"fmt"

	"hookact"
	"hookact/actions/hugo"
)

func main() {
	flag.StringVar(&hugo.RepoLocation, "hugo-repo", "", "Location of the Hugo repository")
	flag.Parse()

	addr := ":8099"
	fmt.Printf("server start at port: %s \n", addr)

	mux := hookact.SetupRoutes()

	err := hookact.StartServer(addr, mux)
	if err != nil {
		fmt.Printf("server start error: %v\n", err)
	}
}
