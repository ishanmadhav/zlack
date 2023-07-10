package main

import "github.com/ishanmadhav/zlack/pkg/api"

func main() {
	tempAPI := api.NewAPI
	tempAPI().Start()
}
