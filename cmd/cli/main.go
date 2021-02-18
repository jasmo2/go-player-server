// CLI
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

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
