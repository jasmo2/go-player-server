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
