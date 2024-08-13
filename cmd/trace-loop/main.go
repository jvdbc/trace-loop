package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type traceloop struct {
	ticker  *time.Ticker
	message string
	begin   uint
}

func newTraceloop(message string, frequency time.Duration, begin uint) *traceloop {
	tl := traceloop{
		ticker:  time.NewTicker(frequency),
		message: message,
		begin:   begin,
	}

	return &tl
}

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "message",
				Value: "trace-loop",
				Usage: "message before counter",
			},
			&cli.DurationFlag{
				Name:  "frequency",
				Value: time.Duration(1 * time.Second),
				Usage: "elasped time between message",
			},
			&cli.UintFlag{
				Name:  "begin",
				Value: 1,
				Usage: "begin counter",
			},
		},
		Name:  "trace-loop",
		Usage: "Display a message with a counter at a certain frequency",
		Action: func(ctx *cli.Context) error {
			msg := ctx.String("message")
			frequency := ctx.Duration("frequency")
			begin := ctx.Uint("begin")

			tl := newTraceloop(msg, frequency, begin)
			tl.run()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func (tl *traceloop) run() {
	count := tl.begin

	for range tl.ticker.C {
		log.Printf("%s: %d", tl.message, count)
		count++
	}
}
