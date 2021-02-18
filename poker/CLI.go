package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI struct
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

// PlayPoker func
func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
