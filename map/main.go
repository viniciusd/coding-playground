package main

import "fmt"

func main() {
	funcs := map[string]func(c map[string]string) int{
		"print": printMap,
	}
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
		"black": "#000000",
	}

	funcs["print"](colors)
}

func printMap(c map[string]string) int {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
	return 42
}
