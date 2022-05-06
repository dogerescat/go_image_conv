package main

import (
	"flag"
	//"fmt"
	"github.com/dogerescat/go_img_conv/search"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	//fmt.Println(dir)
	search.SearchImgFile(dir)
}
