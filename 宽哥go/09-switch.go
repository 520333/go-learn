package main

import "fmt"

func printPriceWithSwitch(weather string) {
	defaultPrice := 10
	switch weather {
	case "lightRain":
		fmt.Println("下小雨,雨伞价格：", defaultPrice+5)
	case "heayRain":
		fmt.Println("下大雨,雨伞价格：", defaultPrice+15)
	case "rainStorm":
		fmt.Println("下暴雨,雨伞价格：", defaultPrice+20)
	case "snowing", "sunny", "other":
		fmt.Println("晴天雪天雨伞价格：", defaultPrice)
	default:
		fmt.Println("这个天气百年一遇,伞送你我先逃命了")
	}
}

// func main() {
// 	printPriceWithSwitch("lightRain")
// 	printPriceWithSwitch("heayRain")
// 	printPriceWithSwitch("rainStorm")
// 	printPriceWithSwitch("snowing")
// 	printPriceWithSwitch("?")
// }
