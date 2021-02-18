package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jasmo2/go-player-server/poker"
)

const dbName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
