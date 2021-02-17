package main

import (
	"encoding/json"
	"io"
)

//FileSystemPlayerStore struct
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague from DB
func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore from FileSystem
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin ,add to win count
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player == nil {
		league = append(league, Player{name, 1})
	} else {
		player.Wins++
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
