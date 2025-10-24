package main

import gotemp "go.template"

func main() {
	_, err := gotemp.New()
	if err != nil {
		panic(err)
	}
}
