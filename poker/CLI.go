package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// CLI struct
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlert
}

// BlindAlert interface
type BlindAlert interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

//NewCLI initialiser
func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlert) CLI {
	return CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		alerter:     alerter,
	}
}

// PlayPoker func
func (cli *CLI) PlayPoker() {
	cli.alerter.ScheduleAlertAt(5*time.Second, 100)
	input := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(input))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
