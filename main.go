package main

import (
	"bufio"
	"os"
	"strconv"
	"time"
	"fmt"
)

func main(){
	selection := []Song{
		{
			title:		"Mettalica - One",
			likes:		0,
			category:	rock,
		},
		{
			title:		"Lenny Kravitz - Fly Away",
			likes:		0,
			category:	rock,
		},
		{
			title:		"Artic Monkeys - R U Mine",
			likes:		0,
			category:	jazz,
		},
		{
			title:		"TOTO - Rosanna",
			likes:		0,
			category:	jazz,
		},
		{
			title:		"Roddy Rich - Box",
			likes:		0,
			category:	pop,
		},
		{
			title:		"R.E.M - Find the River",
			likes:		0,
			category:	pop,
		},
	}

	p := NewPlaylist(selection)

	p.print()

	scanner := bufio.NewScanner(os.Stdin)

	tck := time.NewTicker(50* time.Millisecond)

	for range tck.C{
		next := p.getNextToPlay()
		p.play(next)

		for scanner.Scan(){
			if scanner.Text()=="x"{
				fmt.Println("exiting player ...")
				os.Exit(0)
			}

			choice, _ :=  strconv.ParseInt(scanner.Text(),0,64)
			next.setLiked(choice==1)
			break
		}
		p.sort()
		p.print()
	}
}