package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		doc, err := goquery.NewDocument("https://twitter.com/" + arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, os.Args[0]+":", err)
			os.Exit(1)
		}
		fmt.Println(doc.Find("img.ProfileAvatar-image").First().Attr("src"))
	}
}
