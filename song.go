package main

import "fmt"

type Category	int

const(
	pop Category = iota
	jazz
	rock
)

type Song struct{
	title		string
	category 	Category
	likes		int
}

func (s *Song) print(){
	msg := fmt.Sprintf("%s (%s), likes: %d", s.title, s.category.string(), s.likes)
	fmt.Println(msg)
}

func (s *Song) setLiked(liked bool){
	if liked{
		s.likes++
	}else{
		if s.likes>0{
			s.likes--
		}	
	}
}

func (c Category) string() string{

	switch c{
	case jazz:
		return "jazz";
	case pop:
		return "pop";
	case rock:
		return "rock"
	default:
		return ""	
		
	}
}

