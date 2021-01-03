package main

import "fmt"

type colors map[string]string

func main() {
	colors1 := map[string]string{
		"red":   "#f00",
		"green": "#0f0",
		"blue":  "#00f",
	}
	fmt.Println(colors1)

	colors2 := make(colors)
	colors2["red"] = "#f00"
	colors2["green"] = "#0f0"
	colors2["blue"] = "#00f"
	colors2.print()
}

func (c colors) print() {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
