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
	cli.scheduleBlindAlerts()
	input := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(input))
}

// scheduleBlind func
func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
