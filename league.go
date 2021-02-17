package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League Type
type League []Player

// NewLeague creator
func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

// Find a League by name
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
