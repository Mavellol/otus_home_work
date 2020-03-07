package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05 +0000 UTC"

	now := time.Now()
	ntpNow, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("current time: %s\n", now.Local().Format(layout))
	fmt.Printf("exact time: %s\n", ntpNow.Local().Format(layout))
}
