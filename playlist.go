package main

import (
	"fmt"
	"math/rand"
	"time"
)

//SongList type
type SongList []*Song

//PLaylist type
type Playlist struct{
	songs		SongList
	nowPlaying	*Song
}

func (songs SongList) random() *Song{
	if len(songs) == 0{
		return nil
	}

	//this seed ensure we get a random value evertime it's evaluated
	rand.Seed(time.Now().UnixNano())

	rndPos := rand.Intn(len(songs))

	return songs[rndPos]
}


// New Playlist serves as actor
func NewPlaylist(songs []Song) *Playlist{
	p := Playlist{}
	for i := range songs{
		p.songs = append(p.songs, &songs[i])
	}

	return &p
}

func (p *Playlist) print(){
	fmt.Println("\n Current Playlist is:")


	for _, song := range p.songs{
		song.print()
	}
}


func (p *Playlist) resetLikes(){
	for _,s := range p.songs{
		s.likes = 0
	}
}

func (p *Playlist) play(s *Song){
	p.nowPlaying = s
	fmt.Println(fmt.Sprintf("\nNow playing: %s (%s)\npress 1 / 0 to like / dislike", s.title, s.category.string()))
}

func (p *Playlist) getNextToPlay() *Song{
	var otherCategSongs	SongList

	if p.nowPlaying !=nil{
		for _, s := range p.songs{
			if s.title != p.nowPlaying.title && s.likes < 2 {
				if s.category == p.nowPlaying.category {
					return s
				}
				otherCategSongs = append(otherCategSongs, s)
			}
		}
	}

	// if the current category was listened to twice , try another category
	if len(otherCategSongs) >0 {
		return otherCategSongs.random()
	}

	p.resetLikes()

	return p.songs.random()
}