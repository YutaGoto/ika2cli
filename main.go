package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	"./salmon"
	"./battle"
)

var (
	mode = kingpin.Flag("mode", "Set verbose mode").Short('m').String()
	next = kingpin.Flag("next", "next terms").Short('n').Bool()
	search = kingpin.Flag("search", "search mode").Short('s').Bool()
	league = kingpin.Flag("league", "search next league").Short('l').String()
	gachi = kingpin.Flag("gachi", "search next gachi").Short('g').String()
)

func run(args []string) {
	kingpin.Version("0.4.0")
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Println(err)
	}

	if *mode == "salmon" {
		salmon.GetSalmons(*next)
	} else if *search {
		battle.SearchBattles(*league, *gachi)
	} else {
		battle.GetBattles(*next)
	}

}

func main() {
	run(os.Args[1:])
}
