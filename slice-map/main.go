package main

import "fmt"

func main() {
	/** Slice **/
	// var src []string
	src := []string{"A", "B", "C"}
	src = append(src, "D")
	src = append(src, "E")
	fmt.Println("src:", src)
	/** Copy To Limited-Size Destination & Update Source **/
	dest := make([]string, 3)
	fmt.Println("dest:", dest)
	copy(dest, src)
	fmt.Println("dest:", dest)
	src[0] = "a"
	fmt.Println("src:", src)
	fmt.Println("dest:", dest)
	/** Slice & Update Source **/
	s := src[1:3]
	fmt.Println("s:", s)
	src[2] = "c"
	fmt.Println("src:", src)
	fmt.Println("s:", s)
	/** 2D Slice **/
	gameMap := [][]string{
		make([]string, 3),
		make([]string, 5),
		{"-", "-"},
	}
	fmt.Println("gameMap:", gameMap)
	gameMap[0] = make([]string, 6)
	gameMap[1] = make([]string, 10)
	fmt.Println("gameMap:", gameMap)
	gameMap[1][2] = "-"
	fmt.Println("gameMap:", gameMap)

	/** Map **/
	m := map[string]bool{
		"x": true,
		"y": false,
	}
	fmt.Println("m:", m)
	scores := make(map[string]int)
	fmt.Println("scores:", scores)
	scores["Suthinan"] = 10
	scores["Musitmani"] = 20
	scores["Potaesm"] = 30
	fmt.Println("scores:", scores)
	fmt.Println("Suthinan's score:", scores["Suthinan"])
	delete(scores, "Potaesm")
	value, ok := scores["Potaesm"]
	if ok {
		fmt.Println("Potaesm's score:", value)
	} else {
		fmt.Println("Missing key: Potaesm")
	}
	delete(scores, "UnknownKey")
	fmt.Println("scores:", scores)
	mp := map[string][]string{
		"admins": {"Suthinan", "Musitmani"},
		"users":  {"A", "B", "C"},
	}
	for key, value := range mp {
		fmt.Println("key:", key)
		for index, v := range value {
			fmt.Println("index:", index, "value:", v)
		}
	}
}
