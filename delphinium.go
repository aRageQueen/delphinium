package main

import (
	"delphinium/builder"
	"delphinium/internal"
	"delphinium/output"
	"delphinium/sources"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() //load env variables first (API Keys, etc.,.)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	target := flag.String("target", "", "Domain or organization name")
	format := flag.String("format", "md", "Output format: md")
	outPath := flag.String("out", "timeline.md", "Output file path") // save to file timeline.md by default, or specified file.md
	flag.Parse()

	if *target == "" {
		log.Fatal("Please provide a target using --target")
	}

	var allEvents []internal.Event

	allEvents = append(allEvents, sources.FetchWHOIS(*target)...)
	// allEvents = append(allEvents, sources.FetchBreaches(*target)...)
	// allEvents = append(allEvents, sources.FetchNewsMentions(*target)...)

	timeline := builder.BuildTimeline(allEvents)

	err = output.RenderTimeline(timeline, *format, *outPath)
	if err != nil {
		log.Fatalf("Error rendering output: %v", err)
	}

	fmt.Printf("Timeline for '%s' written to %s\n", *target, *outPath)
}
