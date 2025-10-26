package main

import gotemp "go.template"

func main() {
	g, err := gotemp.New()
	if err != nil {
		panic(err)
	}
	err = g.Open()
	if err != nil {
		panic(err)
	}
}
