package main

import (
	"fmt"
	"os"

	"github.com/ArakiTakaki/scrapbox_cli/dao"
	"github.com/sahilm/fuzzy"
)

func main() {

	exampleFizzy()

	b, err := dao.Say("araki-rust")
	if err != nil {
		os.Exit(0)
	}
	fmt.Println(b.Pages[0].Title)
	for _, v := range b.Pages {
		fmt.Println(v.Title)
	}
	os.Exit(0)
}

func exampleFizzy() {
	const bold = "\033[1m%s\033[0m"
	pattern := "mnres"
	data := []string{"game.cpp", "moduleNameResolver.ts", "my name is_Ramsey"}
	matches := fuzzy.Find(pattern, data)
	for _, match := range matches {
		for i := 0; i < len(match.Str); i++ {
			if contains(i, match.MatchedIndexes) {
				fmt.Print(fmt.Sprintf(bold, string(match.Str[i])))
			} else {
				fmt.Print(string(match.Str[i]))
			}
		}
		fmt.Println()
	}
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}
