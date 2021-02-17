package main

import (
	"io"
)

//FileSystemPlayerStore struct
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague from DB
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore from FileSystem
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}
