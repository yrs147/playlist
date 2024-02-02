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
			title:     "Metallica - Enter Sandman",
			likes:     0,
			category:  rock,
		},
		{
			title:     "Led Zeppelin - Stairway to Heaven",
			likes:     0,
			category:  rock,
		},
		{
			title:     "Miles Davis - So What",
			likes:     0,
			category:  jazz,
		},
		{
			title:     "John Coltrane - Giant Steps",
			likes:     0,
			category:  jazz,
		},
		{
			title:     "Michael Jackson - Billie Jean",
			likes:     0,
			category:  pop,
		},
		{
			title:     "Bruno Mars - Uptown Funk",
			likes:     0,
			category:  pop,
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