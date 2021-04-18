package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("The hexadecimal value to %v is %v \n", color, hex)
	}
}
