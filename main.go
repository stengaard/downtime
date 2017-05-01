// Command downtime calculates availability ratios for a range of dates
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

func usage() {
	me := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", me)
	fmt.Fprintf(os.Stderr, " %s <start_date> <end_date> <downtime_duration>\n", me)
	fmt.Fprintf(os.Stderr, "Note: Both dates are included in the time range")
	flag.PrintDefaults()

}

func main() {
	format := flag.String("date-format", "2006-01-02", "Go time format to parse")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 3 {
		fmt.Fprintf(os.Stderr, "Please supply start date, end date and downtime duration\n")
		os.Exit(1)
	}

	start, err := time.Parse(*format, flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse start date: %v\n", err)
		os.Exit(1)
	}
	end, err := time.Parse(*format, flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse end date: %v", err)
		os.Exit(1)
	}
	// [d1 -- d2], since the date is 00:00 UTC we should really add 24 hours to the
	// end date in order to include that whole day.
	end = end.Add(24 * time.Hour)
	d, err := time.ParseDuration(flag.Arg(2))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse downtime duration: %v", err)
	}

	interval := end.Sub(start)

	fmt.Printf("1-(%v/%v) = %.06f\n", d.Minutes(), interval.Minutes(), float64(1.0-d.Minutes()/interval.Minutes()))

}
