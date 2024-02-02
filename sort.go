package main

// ByLikes custom comparer for songs
type ByLikes Playlist


// Len length of songs
func (b ByLikes) Len() int{
	return len(b.songs)
}

//Swap order of items
func (b ByLikes) Swap(x,y int){
	b.songs[x] , b.songs[y] =  b.songs[y] , b.songs[x]
}
	
func (b ByLikes) Less(x,y int) bool {
	return b.songs[x].likes  > b.songs[y].likes
}	